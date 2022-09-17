// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Judges", testJudges)
	t.Run("Partials", testPartials)
	t.Run("ProblemCategories", testProblemCategories)
	t.Run("ProblemRels", testProblemRels)
	t.Run("ProblemTags", testProblemTags)
	t.Run("Submissions", testSubmissions)
	t.Run("Tags", testTags)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Judges", testJudgesDelete)
	t.Run("Partials", testPartialsDelete)
	t.Run("ProblemCategories", testProblemCategoriesDelete)
	t.Run("ProblemRels", testProblemRelsDelete)
	t.Run("ProblemTags", testProblemTagsDelete)
	t.Run("Submissions", testSubmissionsDelete)
	t.Run("Tags", testTagsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Judges", testJudgesQueryDeleteAll)
	t.Run("Partials", testPartialsQueryDeleteAll)
	t.Run("ProblemCategories", testProblemCategoriesQueryDeleteAll)
	t.Run("ProblemRels", testProblemRelsQueryDeleteAll)
	t.Run("ProblemTags", testProblemTagsQueryDeleteAll)
	t.Run("Submissions", testSubmissionsQueryDeleteAll)
	t.Run("Tags", testTagsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Judges", testJudgesSliceDeleteAll)
	t.Run("Partials", testPartialsSliceDeleteAll)
	t.Run("ProblemCategories", testProblemCategoriesSliceDeleteAll)
	t.Run("ProblemRels", testProblemRelsSliceDeleteAll)
	t.Run("ProblemTags", testProblemTagsSliceDeleteAll)
	t.Run("Submissions", testSubmissionsSliceDeleteAll)
	t.Run("Tags", testTagsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Judges", testJudgesExists)
	t.Run("Partials", testPartialsExists)
	t.Run("ProblemCategories", testProblemCategoriesExists)
	t.Run("ProblemRels", testProblemRelsExists)
	t.Run("ProblemTags", testProblemTagsExists)
	t.Run("Submissions", testSubmissionsExists)
	t.Run("Tags", testTagsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Judges", testJudgesFind)
	t.Run("Partials", testPartialsFind)
	t.Run("ProblemCategories", testProblemCategoriesFind)
	t.Run("ProblemRels", testProblemRelsFind)
	t.Run("ProblemTags", testProblemTagsFind)
	t.Run("Submissions", testSubmissionsFind)
	t.Run("Tags", testTagsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Judges", testJudgesBind)
	t.Run("Partials", testPartialsBind)
	t.Run("ProblemCategories", testProblemCategoriesBind)
	t.Run("ProblemRels", testProblemRelsBind)
	t.Run("ProblemTags", testProblemTagsBind)
	t.Run("Submissions", testSubmissionsBind)
	t.Run("Tags", testTagsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Judges", testJudgesOne)
	t.Run("Partials", testPartialsOne)
	t.Run("ProblemCategories", testProblemCategoriesOne)
	t.Run("ProblemRels", testProblemRelsOne)
	t.Run("ProblemTags", testProblemTagsOne)
	t.Run("Submissions", testSubmissionsOne)
	t.Run("Tags", testTagsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Judges", testJudgesAll)
	t.Run("Partials", testPartialsAll)
	t.Run("ProblemCategories", testProblemCategoriesAll)
	t.Run("ProblemRels", testProblemRelsAll)
	t.Run("ProblemTags", testProblemTagsAll)
	t.Run("Submissions", testSubmissionsAll)
	t.Run("Tags", testTagsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Judges", testJudgesCount)
	t.Run("Partials", testPartialsCount)
	t.Run("ProblemCategories", testProblemCategoriesCount)
	t.Run("ProblemRels", testProblemRelsCount)
	t.Run("ProblemTags", testProblemTagsCount)
	t.Run("Submissions", testSubmissionsCount)
	t.Run("Tags", testTagsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Judges", testJudgesHooks)
	t.Run("Partials", testPartialsHooks)
	t.Run("ProblemCategories", testProblemCategoriesHooks)
	t.Run("ProblemRels", testProblemRelsHooks)
	t.Run("ProblemTags", testProblemTagsHooks)
	t.Run("Submissions", testSubmissionsHooks)
	t.Run("Tags", testTagsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Judges", testJudgesInsert)
	t.Run("Judges", testJudgesInsertWhitelist)
	t.Run("Partials", testPartialsInsert)
	t.Run("Partials", testPartialsInsertWhitelist)
	t.Run("ProblemCategories", testProblemCategoriesInsert)
	t.Run("ProblemCategories", testProblemCategoriesInsertWhitelist)
	t.Run("ProblemRels", testProblemRelsInsert)
	t.Run("ProblemRels", testProblemRelsInsertWhitelist)
	t.Run("ProblemTags", testProblemTagsInsert)
	t.Run("ProblemTags", testProblemTagsInsertWhitelist)
	t.Run("Submissions", testSubmissionsInsert)
	t.Run("Submissions", testSubmissionsInsertWhitelist)
	t.Run("Tags", testTagsInsert)
	t.Run("Tags", testTagsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ProblemCategoryToProblemCategoryUsingParent", testProblemCategoryToOneProblemCategoryUsingParent)
	t.Run("ProblemRelToProblemCategoryUsingCategory", testProblemRelToOneProblemCategoryUsingCategory)
	t.Run("ProblemTagToProblemRelUsingProblem", testProblemTagToOneProblemRelUsingProblem)
	t.Run("ProblemTagToTagUsingTag", testProblemTagToOneTagUsingTag)
	t.Run("ProblemTagToUserUsingUser", testProblemTagToOneUserUsingUser)
	t.Run("SubmissionToUserUsingUser", testSubmissionToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ProblemCategoryToParentProblemCategories", testProblemCategoryToManyParentProblemCategories)
	t.Run("ProblemCategoryToCategoryProblemRels", testProblemCategoryToManyCategoryProblemRels)
	t.Run("ProblemRelToProblemProblemTags", testProblemRelToManyProblemProblemTags)
	t.Run("TagToProblemTags", testTagToManyProblemTags)
	t.Run("UserToProblemTags", testUserToManyProblemTags)
	t.Run("UserToSubmissions", testUserToManySubmissions)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ProblemCategoryToProblemCategoryUsingParentProblemCategories", testProblemCategoryToOneSetOpProblemCategoryUsingParent)
	t.Run("ProblemRelToProblemCategoryUsingCategoryProblemRels", testProblemRelToOneSetOpProblemCategoryUsingCategory)
	t.Run("ProblemTagToProblemRelUsingProblemProblemTags", testProblemTagToOneSetOpProblemRelUsingProblem)
	t.Run("ProblemTagToTagUsingProblemTags", testProblemTagToOneSetOpTagUsingTag)
	t.Run("ProblemTagToUserUsingProblemTags", testProblemTagToOneSetOpUserUsingUser)
	t.Run("SubmissionToUserUsingSubmissions", testSubmissionToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("ProblemCategoryToProblemCategoryUsingParentProblemCategories", testProblemCategoryToOneRemoveOpProblemCategoryUsingParent)
	t.Run("ProblemRelToProblemCategoryUsingCategoryProblemRels", testProblemRelToOneRemoveOpProblemCategoryUsingCategory)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ProblemCategoryToParentProblemCategories", testProblemCategoryToManyAddOpParentProblemCategories)
	t.Run("ProblemCategoryToCategoryProblemRels", testProblemCategoryToManyAddOpCategoryProblemRels)
	t.Run("ProblemRelToProblemProblemTags", testProblemRelToManyAddOpProblemProblemTags)
	t.Run("TagToProblemTags", testTagToManyAddOpProblemTags)
	t.Run("UserToProblemTags", testUserToManyAddOpProblemTags)
	t.Run("UserToSubmissions", testUserToManyAddOpSubmissions)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("ProblemCategoryToParentProblemCategories", testProblemCategoryToManySetOpParentProblemCategories)
	t.Run("ProblemCategoryToCategoryProblemRels", testProblemCategoryToManySetOpCategoryProblemRels)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("ProblemCategoryToParentProblemCategories", testProblemCategoryToManyRemoveOpParentProblemCategories)
	t.Run("ProblemCategoryToCategoryProblemRels", testProblemCategoryToManyRemoveOpCategoryProblemRels)
}

func TestReload(t *testing.T) {
	t.Run("Judges", testJudgesReload)
	t.Run("Partials", testPartialsReload)
	t.Run("ProblemCategories", testProblemCategoriesReload)
	t.Run("ProblemRels", testProblemRelsReload)
	t.Run("ProblemTags", testProblemTagsReload)
	t.Run("Submissions", testSubmissionsReload)
	t.Run("Tags", testTagsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Judges", testJudgesReloadAll)
	t.Run("Partials", testPartialsReloadAll)
	t.Run("ProblemCategories", testProblemCategoriesReloadAll)
	t.Run("ProblemRels", testProblemRelsReloadAll)
	t.Run("ProblemTags", testProblemTagsReloadAll)
	t.Run("Submissions", testSubmissionsReloadAll)
	t.Run("Tags", testTagsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Judges", testJudgesSelect)
	t.Run("Partials", testPartialsSelect)
	t.Run("ProblemCategories", testProblemCategoriesSelect)
	t.Run("ProblemRels", testProblemRelsSelect)
	t.Run("ProblemTags", testProblemTagsSelect)
	t.Run("Submissions", testSubmissionsSelect)
	t.Run("Tags", testTagsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Judges", testJudgesUpdate)
	t.Run("Partials", testPartialsUpdate)
	t.Run("ProblemCategories", testProblemCategoriesUpdate)
	t.Run("ProblemRels", testProblemRelsUpdate)
	t.Run("ProblemTags", testProblemTagsUpdate)
	t.Run("Submissions", testSubmissionsUpdate)
	t.Run("Tags", testTagsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Judges", testJudgesSliceUpdateAll)
	t.Run("Partials", testPartialsSliceUpdateAll)
	t.Run("ProblemCategories", testProblemCategoriesSliceUpdateAll)
	t.Run("ProblemRels", testProblemRelsSliceUpdateAll)
	t.Run("ProblemTags", testProblemTagsSliceUpdateAll)
	t.Run("Submissions", testSubmissionsSliceUpdateAll)
	t.Run("Tags", testTagsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
