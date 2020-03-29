# parser

## Usage

### Teams
```
go run cmd/parser/main.go -mode=teams
```
Will generate `init-teams.sql`

Note: Teams where league.name IN ('Sudamericana', 'Agentes Libres') are skipped.

### Players
```
go run cmd/parser/main.go -mode=players
```
Will generate `init-players.sql`

Note: Players where league.name IN ('Sudamericana', 'Agentes Libres') AND league.team = 'Agentes Libres' are skipped.