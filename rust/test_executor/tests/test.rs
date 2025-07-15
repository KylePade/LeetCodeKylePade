const PROBLEM_FOLDER: &str = "problemsKylePade";
const PROBLEM_ID: &str = "3201";

#[cfg(test)]
mod test {
	use solution_3201 as solution;
    use test_executor::run_test::run_test;

    use crate::{PROBLEM_FOLDER, PROBLEM_ID};

    #[test]
    fn test_solution() {
        run_test(PROBLEM_ID, PROBLEM_FOLDER, solution::solve);
    }
}
