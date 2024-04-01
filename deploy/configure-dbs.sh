#!/bin/bash
# Script to configure PostgreSQL instances for logical replication

max_number_of_replicas=4
max_wal_senders=8
postgres_user="postgres" # Change this to the non-root user that owns the PostgreSQL server process
#synchronous_replication="on"

# Define all PostgreSQL instances
databases=("logrepl_pg_master" "logrepl_pg_replica1" "logrepl_pg_replica2")

for db in "${databases[@]}"; do
    # Check if PostgreSQL is running
    is_running=$(docker exec "$db" su -c "pg_isready" "$postgres_user" | grep "accepting connections")
    if [ -z "$is_running" ]; then
        echo "PostgreSQL is not running on $db. Skipping configuration."
        continue
    fi

    # Apply configuration changes
    docker exec "$db" su -c "sed -i '/#max_replication_slots/s/^#//; /max_replication_slots/s/=.*/= $max_number_of_replicas/' /var/lib/postgresql/data/postgresql.conf" "$postgres_user"
    docker exec "$db" su -c "sed -i '/#max_wal_senders/s/^#//; /max_wal_senders/s/=.*/= $max_wal_senders/' /var/lib/postgresql/data/postgresql.conf"
    #docker exec "$db" su -c "sed -i '/#synchronous_replication/s/^#//; /synchronous_replication/s/=.*/= $synchronous_replication/' /var/lib/postgresql/data/postgresql.conf"

    # Reload PostgreSQL configuration
    docker exec "$db" su -c "pg_ctl reload -D /var/lib/postgresql/data" "$postgres_user"

    # Add replication settings to pg_hba.conf
    docker exec "$db" su -c "echo 'host replication replicator all trust' >> /var/lib/postgresql/data/pg_hba.conf"
    docker exec "$db" su -c "pg_ctl reload -D /var/lib/postgresql/data" "$postgres_user"
done

echo "Configuration applied successfully."
