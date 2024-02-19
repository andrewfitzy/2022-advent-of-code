#!/bin/bash

# Ask for the day
read -p 'Challenge Day: ' day

echo Initialising solution and test folder for day $day

# Copy the three directories and rename to contain the day
cp -R ./day_xx ./day_$day

# Then we'll replace xx in the src and test files with the day number
sed -i '' "s/xx/$day/g" ./day_$day/task_01_main.go
sed -i '' "s/xx/$day/g" ./day_$day/task_02_main.go
sed -i '' "s/xx/$day/g" ./day_$day/task_01_main_test.go
sed -i '' "s/xx/$day/g" ./day_$day/task_02_main_test.go


