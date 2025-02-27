// Copyright 2023 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

import {
  SqlApiResponse,
  executeInternalSql,
  formatApiResult,
  LARGE_RESULT_SIZE,
  LONG_TIMEOUT,
  SqlExecutionRequest,
  sqlResultsAreEmpty,
  SqlTxnResult,
} from "./sqlApi";
import {
  ContentionDetails,
  getInsightsFromProblemsAndCauses,
  InsightExecEnum,
  StmtInsightEvent,
} from "src/insights";
import moment from "moment";
import { INTERNAL_APP_NAME_PREFIX } from "src/recentExecutions/recentStatementUtils";
import { FixFingerprintHexValue } from "../util";
import { getContentionDetailsApi } from "./contentionApi";

export type StmtInsightsReq = {
  start?: moment.Moment;
  end?: moment.Moment;
  stmtExecutionID?: string;
  stmtFingerprintId?: string;
};

export type StmtInsightsResponseRow = {
  session_id: string;
  txn_id: string;
  txn_fingerprint_id: string; // hex string
  implicit_txn: boolean;
  stmt_id: string;
  stmt_fingerprint_id: string; // hex string
  query: string;
  start_time: string; // Timestamp
  end_time: string; // Timestamp
  full_scan: boolean;
  user_name: string;
  app_name: string;
  database_name: string;
  rows_read: number;
  rows_written: number;
  priority: string;
  retries: number;
  exec_node_ids: number[];
  contention: string; // interval
  contention_events: ContentionDetails[];
  last_retry_reason?: string;
  causes: string[];
  problem: string;
  index_recommendations: string[];
  plan_gist: string;
  cpu_sql_nanos: number;
};

const stmtColumns = `
session_id,
txn_id,
txn_fingerprint_id,
implicit_txn,
stmt_id,
stmt_fingerprint_id,
query,
start_time,
end_time,
full_scan,
user_name,
app_name,
database_name,
rows_read,
rows_written,
priority,
retries,
exec_node_ids,
contention,
last_retry_reason,
causes,
problem,
index_recommendations,
plan_gist,
cpu_sql_nanos
`;

const stmtInsightsOverviewQuery = (filters?: StmtInsightsReq): string => {
  if (filters?.stmtExecutionID) {
    return `
SELECT ${stmtColumns} FROM crdb_internal.cluster_execution_insights
WHERE stmt_id = '${filters.stmtExecutionID}'`;
  }

  let whereClause = `
  WHERE app_name NOT LIKE '${INTERNAL_APP_NAME_PREFIX}%'
  AND problem != 'None'
  AND txn_id != '00000000-0000-0000-0000-000000000000'`;
  if (filters?.start) {
    whereClause =
      whereClause + ` AND start_time >= '${filters.start.toISOString()}'`;
  }
  if (filters?.end) {
    whereClause =
      whereClause + ` AND end_time <= '${filters.end.toISOString()}'`;
  }
  if (filters?.stmtFingerprintId) {
    whereClause =
      whereClause +
      ` AND encode(stmt_fingerprint_id, 'hex') = '${filters.stmtFingerprintId}'`;
  }

  return `
SELECT ${stmtColumns} FROM
   (
     SELECT DISTINCT ON (stmt_fingerprint_id, problem, causes)
       *
     FROM
       crdb_internal.cluster_execution_insights
         ${whereClause}
     ORDER BY stmt_fingerprint_id, problem, causes, end_time DESC
   )`;
};

export const stmtInsightsByTxnExecutionQuery = (id: string): string => `
 SELECT ${stmtColumns}
 FROM crdb_internal.cluster_execution_insights
 WHERE txn_id = '${id}'
`;

export async function getStmtInsightsApi(
  req?: StmtInsightsReq,
): Promise<SqlApiResponse<StmtInsightEvent[]>> {
  const request: SqlExecutionRequest = {
    statements: [
      {
        sql: stmtInsightsOverviewQuery(req),
      },
    ],
    execute: true,
    max_result_size: LARGE_RESULT_SIZE,
    timeout: LONG_TIMEOUT,
  };

  const result = await executeInternalSql<StmtInsightsResponseRow>(request);

  if (sqlResultsAreEmpty(result)) {
    return formatApiResult([], result.error, "retrieving insights information");
  }
  const stmtInsightEvent = formatStmtInsights(result.execution?.txn_results[0]);
  await addStmtContentionInfoApi(stmtInsightEvent);
  return formatApiResult(
    stmtInsightEvent,
    result.error,
    "retrieving insights information",
  );
}

async function addStmtContentionInfoApi(
  input: StmtInsightEvent[],
): Promise<void> {
  if (!input || input.length === 0) {
    return;
  }

  for (let i = 0; i < input.length; i++) {
    const event = input[i];
    if (
      event.contentionTime == null ||
      event.contentionTime.asMilliseconds() <= 0
    ) {
      continue;
    }

    event.contentionEvents = await getContentionDetailsApi({
      waitingTxnID: null,
      waitingStmtID: event.statementExecutionID,
      start: null,
      end: null,
    });
  }
}

export function formatStmtInsights(
  response: SqlTxnResult<StmtInsightsResponseRow>,
): StmtInsightEvent[] {
  if (!response?.rows?.length) {
    return [];
  }

  return response.rows.map((row: StmtInsightsResponseRow) => {
    const start = moment.utc(row.start_time);
    const end = moment.utc(row.end_time);

    return {
      transactionExecutionID: row.txn_id,
      transactionFingerprintID: FixFingerprintHexValue(row.txn_fingerprint_id),
      implicitTxn: row.implicit_txn,
      databaseName: row.database_name,
      application: row.app_name,
      username: row.user_name,
      sessionID: row.session_id,
      priority: row.priority,
      retries: row.retries,
      lastRetryReason: row.last_retry_reason,
      query: row.query,
      startTime: start,
      endTime: end,
      elapsedTimeMillis: end.diff(start, "milliseconds"),
      statementExecutionID: row.stmt_id,
      statementFingerprintID: FixFingerprintHexValue(row.stmt_fingerprint_id),
      isFullScan: row.full_scan,
      rowsRead: row.rows_read,
      rowsWritten: row.rows_written,
      // This is the total stmt contention.
      contentionTime: row.contention ? moment.duration(row.contention) : null,
      indexRecommendations: row.index_recommendations,
      insights: getInsightsFromProblemsAndCauses(
        [row.problem],
        row.causes,
        InsightExecEnum.STATEMENT,
      ),
      planGist: row.plan_gist,
      cpuSQLNanos: row.cpu_sql_nanos,
    } as StmtInsightEvent;
  });
}
