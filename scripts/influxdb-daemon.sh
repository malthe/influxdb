#!/usr/bin/env bash

nohup sh -c "/usr/bin/influxdb $@ | logger -p daemon.info -t influxdb" &
