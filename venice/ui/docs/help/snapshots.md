---
id: snapshots
---

# Snapshots

Within the PSM configuration, it is capable to take a snapshot of the current configuration of the PSM environment.  The user is also able to restore any of the snapshots that has been taken.

## Save a Config Snapshot
To create a snapshot, frome the side menu, go to "Admin" and select "Snapshots".  On the top right, click on "Save a Config Snapshot".  PSM will create a snapshot file called "snapshot-xxxxx" with the date appended.  You can download this file for offline storage repository.

## Restore a Snapshot And Upload Snapshot File
To restore a snapshot configuration file, there are 2 options.  Within the UI, a list of snapshot files are shown under "Configurations".  To restore a particular snapshot, on the right side of that file is an restore icon.  Click on that "restore" icon and a pop-up will request you to confirm if you want to restore this snapshot.

In case a snapshot file is not shown on the list, the other option is to upload the snapshot file from your local repository to the PSM controller.  To do this, expand "Upload Snapshot File".  You can then click on "Choose" to select the snapshot file.  Once it is selected, you can click on "Upload" to upload the configuration.  Once it is shown on the UI, you can restore the snapshot by clicking on the "restore" icon from that snapshot.