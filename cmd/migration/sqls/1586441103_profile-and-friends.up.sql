BEGIN;

/**
 * Table `friend`
 *
 * Friends in the application
 */
CREATE TABLE generator.friend(
  id                uuid PRIMARY KEY DEFAULT pgcrypto.gen_random_uuid(),
  sender            uuid NOT NULL REFERENCES generator.user(id),
  receiver          uuid NOT NULL REFERENCES generator.user(id),
  state             text NOT NULL,
  created_at        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        TIMESTAMP NOT NULL
);

ALTER TABLE generator.friend
  ADD CONSTRAINT uniq_friend UNIQUE (sender, receiver)
;

CREATE TRIGGER friend_updated_at
  BEFORE INSERT OR UPDATE ON generator.friend
  FOR EACH ROW
  EXECUTE PROCEDURE generator.set_updated_at()
;

COMMIT;