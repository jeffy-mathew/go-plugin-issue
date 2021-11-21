This repository tries to reproduce `plugin already loaded` issue.
When the same plugin is build and loaded once again from the driver program,
it results in a `plugin already loaded error`. 

Step to follow,

    ./setup.sh
    cd driver
    ./main