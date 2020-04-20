BEGIN;

CREATE SCHEMA IF NOT EXISTS generator;
CREATE SCHEMA IF NOT EXISTS pgcrypto;

CREATE EXTENSION IF NOT EXISTS pgcrypto SCHEMA pgcrypto;

CREATE OR REPLACE FUNCTION generator.set_updated_at()
RETURNS TRIGGER STABLE AS $plpgsql$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$plpgsql$ LANGUAGE plpgsql;

/**
 * Table `fifa_league`
 *
 * League set of FIFA
 */
CREATE TABLE generator.fifa_league(
  id                uuid PRIMARY KEY DEFAULT pgcrypto.gen_random_uuid(),
  name              text NOT NULL
);

ALTER TABLE generator.fifa_league
  ADD CONSTRAINT uniq_fifa_league UNIQUE (name)
;

/**
 * Table `fifa_team`
 *
 * Team set of FIFA
 */
CREATE TABLE generator.fifa_team(
  id                uuid PRIMARY KEY DEFAULT pgcrypto.gen_random_uuid(),
  name              text NOT NULL,
  shield_src        text DEFAULT NULL,
  league_id         uuid NOT NULL REFERENCES generator.fifa_league(id)
);

ALTER TABLE generator.fifa_team
  ADD CONSTRAINT uniq_fifa_team UNIQUE (name, league_id)
;

/**
 * Table `fifa_player`
 *
 * Player set of FIFA
 */
CREATE TABLE generator.fifa_player(
  id                uuid PRIMARY KEY DEFAULT pgcrypto.gen_random_uuid(),
  name              text NOT NULL,
  position          text NOT NULL,
  number            integer NOT NULL,
  picture_src       text DEFAULT NULL,
  team_id           uuid NOT NULL REFERENCES generator.fifa_team(id)
);

-- TODO: CREATE THIS IF WE FIND HOW TO REMOVE THE DUPLICATES
-- ALTER TABLE generator.fifa_player
--   ADD CONSTRAINT uniq_fifa_player UNIQUE (name, team_id)
-- ;

/**
 * Table `user`
 *
 * Users registered to our platform
 */
CREATE TABLE generator.user(
  id                uuid PRIMARY KEY DEFAULT pgcrypto.gen_random_uuid(),
  name              text NOT NULL,
  username          text NOT NULL,
  password          text NOT NULL,
  profile_picture   text DEFAULT NULL,
  active            bool NOT NULL DEFAULT FALSE,
  created_at        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        TIMESTAMP NOT NULL
);

ALTER TABLE generator.user
  ADD CONSTRAINT uniq_user UNIQUE (username)
;

/**
 * Table `group`
 *
 * Group information
 */
CREATE TABLE generator.group(
  id                uuid PRIMARY KEY DEFAULT pgcrypto.gen_random_uuid(),
  name              text NOT NULL,
  created_at        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        TIMESTAMP NOT NULL
);

/**
 * Table `user_group`
 *
 * Set of Users in a Group
 */
CREATE TABLE generator.user_group(
  id                uuid PRIMARY KEY DEFAULT pgcrypto.gen_random_uuid(),
  user_id           uuid NOT NULL REFERENCES generator.user(id),
  group_id          uuid NOT NULL REFERENCES generator.group(id),
  created_at        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        TIMESTAMP NOT NULL
);

/**
 * Table `tournament`
 *
 * Tournament information inside a Group
 */
CREATE TABLE generator.tournament(
  id                    uuid PRIMARY KEY DEFAULT pgcrypto.gen_random_uuid(),
  name                  text NOT NULL,
  type                  text NOT NULL,
  num_players           int NOT NULL,
  num_teams_player      int NOT NULL DEFAULT 1,
  times_against         int DEFAULT NULL, /* TODO MIRAR SI PONER AQUI O NO (solo valido para ligas)*/
  round                 text DEFAULT NULL, /* TODO MIRAR SI PONER AQUI O NO (solo valido para copa)*/
  group_id         uuid NOT NULL REFERENCES generator.group(id),
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL
);

/**
 * Table `user_tournament`
 *
 * Relation between User and Tournament
 */
CREATE TABLE generator.user_tournament(
  id                    uuid PRIMARY KEY DEFAULT pgcrypto.gen_random_uuid(),
  user_id               uuid NOT NULL REFERENCES generator.user(id),
  tournament_id         uuid NOT NULL REFERENCES generator.tournament(id),
  team_id               uuid NOT NULL REFERENCES generator.fifa_team(id),
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL
);

/**
 * Table `fixture`
 *
 * Fixtures of the Tournament
 */
CREATE TABLE generator.fixture(
  id                    uuid PRIMARY KEY DEFAULT pgcrypto.gen_random_uuid(),
  name                  text NOT NULL,
  tournament_id         uuid NOT NULL REFERENCES generator.tournament(id),
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL
);

/**
 * Table `match`
 *
 * Matches in a Fixture
 */
CREATE TABLE generator.match(
  id                    uuid PRIMARY KEY DEFAULT pgcrypto.gen_random_uuid(),
  fixture_id            uuid NOT NULL REFERENCES generator.fixture(id),
  home                  uuid NOT NULL REFERENCES generator.user_tournament(id),
  away                  uuid NOT NULL REFERENCES generator.user_tournament(id),
  played                bool NOT NULL DEFAULT FALSE,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL
);

/**
 * Table `goal`
 *
 * Goals in a Match
 */
CREATE TABLE generator.goal(
  id                    uuid PRIMARY KEY DEFAULT pgcrypto.gen_random_uuid(),
  match_id              uuid NOT NULL REFERENCES generator.match(id),
  player_id             uuid NOT NULL REFERENCES generator.fifa_player(id),
  type                  text DEFAULT NULL,
  minute                integer DEFAULT NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/**
 * Triggers
 */
CREATE TRIGGER user_updated_at
  BEFORE INSERT OR UPDATE ON generator.user
  FOR EACH ROW
  EXECUTE PROCEDURE generator.set_updated_at()
;

CREATE TRIGGER group_updated_at
  BEFORE INSERT OR UPDATE ON generator.group
  FOR EACH ROW
  EXECUTE PROCEDURE generator.set_updated_at()
;

CREATE TRIGGER tournament_updated_at
  BEFORE INSERT OR UPDATE ON generator.tournament
  FOR EACH ROW
  EXECUTE PROCEDURE generator.set_updated_at()
;

CREATE TRIGGER user_tournament_updated_at
  BEFORE INSERT OR UPDATE ON generator.user_tournament
  FOR EACH ROW
  EXECUTE PROCEDURE generator.set_updated_at()
;

CREATE TRIGGER fixture_updated_at
  BEFORE INSERT OR UPDATE ON generator.fixture
  FOR EACH ROW
  EXECUTE PROCEDURE generator.set_updated_at()
;

CREATE TRIGGER match_updated_at
  BEFORE INSERT OR UPDATE ON generator.match
  FOR EACH ROW
  EXECUTE PROCEDURE generator.set_updated_at()
;

COMMIT;
