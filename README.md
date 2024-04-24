## DB Migrate
Replace `project_path` according to your needs.

### Migrate up
```shell
migrate -path 'repository/sql/migrations' -database 'postgres://<YOUR-NAME>:<YOUR-PASSWORD>@aws-0-eu-central-1.pooler.supabase.com:5432/postgres' up
```

### Migrate down
```shell
migrate -path 'repository/sql/migrations' -database 'postgres://<YOUR-NAME>:<YOUR-PASSWORD>@aws-0-eu-central-1.pooler.supabase.com:5432/postgres' down
```