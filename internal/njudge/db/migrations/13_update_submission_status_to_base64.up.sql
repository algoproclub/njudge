-- This migration used to convert the free-text fields of submissions.status to base64 with a Go function.
-- All databases are past it, so it is a no-op now; the version has to stay for golang-migrate's sequence.
SELECT 1;
