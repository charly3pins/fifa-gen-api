BEGIN;

/**
 * Table `friendship`
 *
 * Friends in the application
 */
CREATE TABLE generator.friendship(
  user_one_id       uuid NOT NULL REFERENCES generator.user(id),
  user_two_id       uuid NOT NULL REFERENCES generator.user(id),
  status            int NOT NULL,
  action_user_id    uuid NOT NULL REFERENCES generator.user(id),
  created_at        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        TIMESTAMP NOT NULL
);

/**
STATUS
Code	Meaning
0	    Pending
1	    Accepted
2	    Declined
3	    Blocked
*/

ALTER TABLE generator.friendship
  ADD CONSTRAINT uniq_friend UNIQUE (user_one_id, user_two_id)
;

CREATE TRIGGER friendship_updated_at
  BEFORE INSERT OR UPDATE ON generator.friendship
  FOR EACH ROW
  EXECUTE PROCEDURE generator.set_updated_at()
;

COMMIT;
