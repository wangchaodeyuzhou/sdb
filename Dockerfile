FROM sdb as builder

COPY sdb sdb

COPY configs/config.yml configs/config.yml

ENTRYPOINT "./sdb" $0 $@