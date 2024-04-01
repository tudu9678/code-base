#!/bin/bash
# Define the database containers
databases=("logrepl_pg_master" "logrepl_pg_replica1" "logrepl_pg_replica2")
schema_file="./integrations/schema/schema.sql"

for db in "${databases[@]}"; do
  # Check if container is running
  if docker ps -q -f name="$db" >/dev/null 2>&1; then
    # Get mapped port (assuming format "host_port:container_port")
    port=$(docker inspect -f '{{(index (index .NetworkSettings.Ports "5432/tcp") 0).HostPort}}' "$db")

    echo "Copying schema to $db... (port: $port)"
    docker cp "$schema_file" "$db:/tmp/schema.sql"

    echo "Loading schema in $db..."
    # Consider adding error handling for psql execution
    docker exec -i "$db" psql -U postgres -d postgres -f /tmp/schema.sql || true

    echo "Removing schema file from $db..."
    docker exec -i "$db" rm /tmp/schema.sql

    echo "Schema loaded successfully in $db!"
  else
    echo "Skipping $db: Container not running."
  fi
done
