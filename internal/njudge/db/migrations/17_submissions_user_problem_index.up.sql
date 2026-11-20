create index concurrently if not exists submissions_user_id_problem_id_idx
    on public.submissions (user_id, problem_id);
