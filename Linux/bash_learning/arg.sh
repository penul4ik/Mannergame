#!/usr/bin/env bash

starttime=$(awk '{print $22}' /proc/153403/stat)
starttime_in_seconds=$(echo "$starttime / 100" | bc)
echo "starttime - $starttime_in_seconds"
systemtime=$(cut -d' ' -f1 /proc/uptime)
proctime=$(echo "$systemtime - $starttime_in_seconds" | bc)
echo $proctime
proctime_hours=$(echo "scale=2; $proctime / 3600" | bc)
echo "Process uptime in hours: $proctime_hours"
