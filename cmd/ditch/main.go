package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var skipThese = []string{
	`Snapt\Nova\Adc\Jobs\FetchAdcStats`,
	`Snapt\Nova\Adc\Jobs\FetchAdcStats`,
	`Snapt\Nova\Adc\Jobs\CheckWAF`,
	`Snapt\Nova\Prometheus\Jobs\LoadPrometheusStats`,
	`Snapt\Nova\Jobs\JobQueueServerStats`,
	`Snapt\Nova\Announcements\Jobs\QueueAnnouncementEmails`,
	`Snapt\Nova\Waf\Jobs\WafLogs`,
	`Snapt\Nova\Jobs\JobQueueServerStats`,
	`Snapt\Nova\Jobs\JobQueueProcessAlerts`,
	`Snapt\Nova\Prometheus\Jobs\OrgPrometheusStats`,
	`Snapt\Nova\Prometheus\Jobs\NovaPrometheusStats`,
	`Snapt\Nova\Nodes\Snmp\Jobs\DeploySnmpToNode`,
	`Snapt\Nova\Listeners\Adcs\NodeUpdater`,
	`Snapt\Nova\Nodes\NodeShield\Jobs\DeployShieldToNode`,

	`reusable`,
	`epoll`,
	`timer delta:`,
	`free:`,
	`worker cycle`,
	`event timer`,
	`shmtx`,
}

func skip(line string) bool {
	for _, skipThis := range skipThese {
		if strings.Contains(line, skipThis) {
			if !strings.Contains(line, "Failed") {
				return true
			}
		}
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		if skip(line) {
			continue
		}

		fmt.Print(line)
	}
}
