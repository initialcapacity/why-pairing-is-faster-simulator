# Why pairing is faster simulator

Simulate pairs and solo developers working on a backlog.
Pairs complete stories immediately through live code review, while
solo developers must send code through a PR and revision process.
Pairs are also able to review code written by solo developers.

Read the [paper](https://github.com/initialcapacity/why-pairing-is-faster).

## Running the simulator

Run the simulation with the below command. Modify `simulator.go` for both pairing,
individual and mixed teams.

```shell
go run cmd/simulator/simulator.go
```

## Example pairing run

Simulates 2 pairs working on a backlog of 19 stories with dependencies.
The backlog is completed in 34 hours with 38 points of communication.

```
Starting work
Pair 'A' starting story 'Create person'
Pair 'A' finished story 'Create person'
Pair 'A' starting story 'Create project'
Pair 'B' starting story 'List people'
Pair 'B' finished story 'List people'
Pair 'B' starting story 'Create timesheet'
Pair 'A' finished story 'Create project'
Pair 'A' starting story 'Show person'
Pair 'A' finished story 'Show person'
Pair 'A' starting story 'Edit person'
Pair 'B' finished story 'Create timesheet'
Pair 'B' starting story 'List projects'
Pair 'B' finished story 'List projects'
Pair 'B' starting story 'List timesheets'
Pair 'A' finished story 'Edit person'
Pair 'A' starting story 'Show project'
Pair 'A' finished story 'Show project'
Pair 'A' starting story 'Edit project'
Pair 'B' finished story 'List timesheets'
Pair 'B' starting story 'Show timesheet'
Pair 'B' finished story 'Show timesheet'
Pair 'B' starting story 'Delete timesheet'
Pair 'A' finished story 'Edit project'
Pair 'A' starting story 'Edit timesheet'
Pair 'B' finished story 'Delete timesheet'
Pair 'B' starting story 'Assign person to timesheet'
Pair 'A' finished story 'Edit timesheet'
Pair 'A' starting story 'Assign project to timesheet'
Pair 'B' finished story 'Assign person to timesheet'
Pair 'B' starting story 'Submit timesheet'
Pair 'A' finished story 'Assign project to timesheet'
Pair 'B' finished story 'Submit timesheet'
Pair 'A' starting story 'Review timesheet'
Pair 'A' finished story 'Review timesheet'
Pair 'B' starting story 'Accept timesheet'
Pair 'A' starting story 'Reject timeheet'
Pair 'B' finished story 'Accept timesheet'
Pair 'A' finished story 'Reject timeheet'
Backlog complete in 34 hours
```

## Example pull-request run

Simulates 4 individuals working on a backlog of 19 stories with dependencies.
Backlog is completed in 43 hours with 114 points of communication. 

```
Starting work
Solo 'A' starting story 'Create person'
Solo 'A' finished story 'Create person'
Solo 'B' reviewing story 'Create person'
Solo 'B' reviewed story 'Create person'
Solo 'A' resolving story 'Create person'
Solo 'A' resolved story 'Create person'
Solo 'B' starting story 'List people'
Solo 'C' starting story 'Create project'
Solo 'D' starting story 'Create timesheet'
Solo 'B' finished story 'List people'
Solo 'A' reviewing story 'List people'
Solo 'C' finished story 'Create project'
Solo 'D' finished story 'Create timesheet'
Solo 'A' reviewed story 'List people'
Solo 'B' reviewing story 'Create project'
Solo 'C' reviewing story 'Create timesheet'
Solo 'B' reviewed story 'Create project'
Solo 'C' reviewed story 'Create timesheet'
Solo 'B' resolving story 'List people'
Solo 'C' resolving story 'Create project'
Solo 'D' resolving story 'Create timesheet'
Solo 'B' resolved story 'List people'
Solo 'C' resolved story 'Create project'
Solo 'D' resolved story 'Create timesheet'
Solo 'A' starting story 'Show person'
Solo 'B' starting story 'Edit person'
Solo 'D' starting story 'List projects'
Solo 'C' starting story 'List timesheets'
Solo 'A' finished story 'Show person'
Solo 'D' finished story 'List projects'
Solo 'C' finished story 'List timesheets'
Solo 'A' reviewing story 'List projects'
Solo 'D' reviewing story 'List timesheets'
Solo 'C' reviewing story 'Show person'
Solo 'B' finished story 'Edit person'
Solo 'A' reviewed story 'List projects'
Solo 'D' reviewed story 'List timesheets'
Solo 'A' reviewing story 'Edit person'
Solo 'C' reviewed story 'Show person'
Solo 'A' reviewed story 'Edit person'
Solo 'D' resolving story 'List projects'
Solo 'B' resolving story 'Edit person'
Solo 'D' resolved story 'List projects'
Solo 'A' starting story 'Show project'
Solo 'D' starting story 'Edit project'
Solo 'C' resolving story 'List timesheets'
Solo 'B' resolved story 'Edit person'
Solo 'C' resolved story 'List timesheets'
Solo 'C' starting story 'Show timesheet'
Solo 'B' starting story 'Delete timesheet'
Solo 'A' finished story 'Show project'
Solo 'A' starting story 'Edit timesheet'
Solo 'B' finished story 'Delete timesheet'
Solo 'B' starting story 'Assign person to timesheet'
Solo 'C' finished story 'Show timesheet'
Solo 'C' starting story 'Assign project to timesheet'
Solo 'D' finished story 'Edit project'
Solo 'D' reviewing story 'Show project'
Solo 'D' reviewed story 'Show project'
Solo 'D' reviewing story 'Show timesheet'
Solo 'A' finished story 'Edit timesheet'
Solo 'A' reviewing story 'Delete timesheet'
Solo 'D' reviewed story 'Show timesheet'
Solo 'D' reviewing story 'Edit timesheet'
Solo 'C' finished story 'Assign project to timesheet'
Solo 'B' finished story 'Assign person to timesheet'
Solo 'B' reviewing story 'Edit project'
Solo 'C' resolving story 'Show timesheet'
Solo 'A' reviewed story 'Delete timesheet'
Solo 'A' reviewing story 'Assign person to timesheet'
Solo 'D' reviewed story 'Edit timesheet'
Solo 'D' reviewing story 'Assign project to timesheet'
Solo 'D' reviewed story 'Assign project to timesheet'
Solo 'B' reviewed story 'Edit project'
Solo 'A' reviewed story 'Assign person to timesheet'
Solo 'C' resolved story 'Show timesheet'
Solo 'A' resolving story 'Show person'
Solo 'D' resolving story 'Edit project'
Solo 'B' resolving story 'Delete timesheet'
Solo 'C' resolving story 'Assign project to timesheet'
Solo 'C' resolved story 'Assign project to timesheet'
Solo 'A' resolved story 'Show person'
Solo 'D' resolved story 'Edit project'
Solo 'B' resolved story 'Delete timesheet'
Solo 'A' resolving story 'Show project'
Solo 'B' resolving story 'Assign person to timesheet'
Solo 'A' resolved story 'Show project'
Solo 'B' resolved story 'Assign person to timesheet'
Solo 'A' starting story 'Submit timesheet'
Solo 'A' finished story 'Submit timesheet'
Solo 'D' reviewing story 'Submit timesheet'
Solo 'A' resolving story 'Edit timesheet'
Solo 'D' reviewed story 'Submit timesheet'
Solo 'A' resolved story 'Edit timesheet'
Solo 'A' resolving story 'Submit timesheet'
Solo 'A' resolved story 'Submit timesheet'
Solo 'D' starting story 'Review timesheet'
Solo 'D' finished story 'Review timesheet'
Solo 'C' reviewing story 'Review timesheet'
Solo 'C' reviewed story 'Review timesheet'
Solo 'D' resolving story 'Review timesheet'
Solo 'D' resolved story 'Review timesheet'
Solo 'C' starting story 'Accept timesheet'
Solo 'B' starting story 'Reject timeheet'
Solo 'B' finished story 'Reject timeheet'
Solo 'C' finished story 'Accept timesheet'
Solo 'A' reviewing story 'Accept timesheet'
Solo 'D' reviewing story 'Reject timeheet'
Solo 'A' reviewed story 'Accept timesheet'
Solo 'D' reviewed story 'Reject timeheet'
Solo 'C' resolving story 'Accept timesheet'
Solo 'B' resolving story 'Reject timeheet'
Solo 'B' resolved story 'Reject timeheet'
Solo 'C' resolved story 'Accept timesheet'
Backlog complete in 43 hours
```
