#!/bin/bash
/root/messenger/messenger dashboard &
/root/messenger/messenger jobs notification -n 2 &
/root/messenger/messenger jobs sender --type mail -n 2
