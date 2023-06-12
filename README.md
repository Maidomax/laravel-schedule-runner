# Laravel Schedule Runner

A small Golang program to replace the cron for running the Laravel schedule in environments where this is not possible (like Digital Ocean App Platform).

## Instructions

Copy over the `laravel-schedule-runner` binary into the root of your Laravel project. Before the command that starts your app, run `nohup ./laravel-schedule-runner >> /dev/null 2>&1 &`. This will start the schedule runner in the background, and monitor the process.

In environments where consecutive deployments may happen in-place on the same container or VPS, you can create a bash script like this:

```
#!/bin/bash

# Schedule runner
if [ "$(/usr/bin/pgrep -f "laravel-schedule-runner")" ]; then
    echo "Schedule runner already working"
else
    echo "Initializing schedule runner"
    nohup ./laravel-schedule-runner >> /dev/null 2>&1 &
fi
```

And use that to start the runner. This will make sure only one instance is running.

In the local environment, you can just start the runner with:
```
./laravel-schedule-runner
```