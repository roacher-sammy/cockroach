// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	cluster2 "github.com/cockroachdb/cockroach/pkg/cmd/roachtest/cluster"
	"github.com/cockroachdb/cockroach/pkg/cmd/roachtest/spec"
)

func registerPebble(r *testRegistry) {
	pebble := os.Getenv("PEBBLE_BIN")
	if pebble == "" {
		pebble = "./pebble.linux"
	}

	run := func(ctx context.Context, t *test, c cluster2.Cluster, size int) {
		c.Put(ctx, pebble, "./pebble")

		const initialKeys = 10_000_000
		const cache = 4 << 30 // 4 GB
		const duration = 10 * time.Minute
		const dataDir = "$(dirname {store-dir})"
		const dataTar = dataDir + "/data.tar"
		const benchDir = dataDir + "/bench"

		runCmd := func(cmd string) {
			t.l.PrintfCtx(ctx, "> %s", cmd)
			err := c.RunL(ctx, t.l, c.All(), cmd)
			t.l.Printf("> result: %+v", err)
			if err := ctx.Err(); err != nil {
				t.l.Printf("(note: incoming context was canceled: %s", err)
			}
			if err != nil {
				t.Fatal(err)
			}
		}

		// Generate the initial DB state. This is somewhat time consuming for
		// larger value sizes, so we do this once and reuse the same DB state on
		// all of the workloads.
		runCmd(fmt.Sprintf(
			"(./pebble bench ycsb %s"+
				" --wipe "+
				" --workload=read=100"+
				" --concurrency=1"+
				" --values=%d"+
				" --initial-keys=%d"+
				" --cache=%d"+
				" --num-ops=1 && "+
				"rm -f %s && tar cvPf %s %s) > init.log 2>&1",
			benchDir, size, initialKeys, cache, dataTar, dataTar, benchDir))

		for _, workload := range []string{"A", "B", "C", "D", "E", "F"} {
			keys := "zipf"
			switch workload {
			case "D":
				keys = "uniform"
			}

			runCmd(fmt.Sprintf(
				"rm -fr %s && tar xPf %s &&"+
					" ./pebble bench ycsb %s"+
					" --workload=%s"+
					" --concurrency=256"+
					" --values=%d"+
					" --keys=%s"+
					" --initial-keys=0"+
					" --prepopulated-keys=%d"+
					" --cache=%d"+
					" --duration=%s > ycsb.log 2>&1",
				benchDir, dataTar, benchDir, workload, size, keys, initialKeys, cache, duration))

			runCmd(fmt.Sprintf("tar cvPf profiles_%s.tar *.prof", workload))

			dest := filepath.Join(t.artifactsDir, fmt.Sprintf("ycsb_%s.log", workload))
			if err := c.Get(ctx, t.l, "ycsb.log", dest, c.All()); err != nil {
				t.Fatal(err)
			}

			profilesName := fmt.Sprintf("profiles_%s.tar", workload)
			dest = filepath.Join(t.artifactsDir, profilesName)
			if err := c.Get(ctx, t.l, profilesName, dest, c.All()); err != nil {
				t.Fatal(err)
			}

			runCmd("rm -fr *.prof")
		}
	}

	for _, size := range []int{64, 1024} {
		size := size
		r.Add(testSpec{
			Name:       fmt.Sprintf("pebble/ycsb/size=%d", size),
			Owner:      OwnerStorage,
			Timeout:    2 * time.Hour,
			MinVersion: "v20.1.0",
			Cluster:    r.makeClusterSpec(5, spec.CPU(16)),
			Tags:       []string{"pebble"},
			Run: func(ctx context.Context, t *test, c cluster2.Cluster) {
				run(ctx, t, c, size)
			},
		})
	}
}
