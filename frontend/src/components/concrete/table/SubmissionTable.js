import { useState } from "react"
import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { SVGSpinner, SVGView } from "../../svg/SVGs"
import RoundedTable from "../../container/RoundedTable"
import CopyableCode from "../../util/copy/CopyableCode"
import Modal from "../../container/modal/Modal"

function TestCase13({ index, numCases, testCase, group, isLastGroup, isLastCase }) {
    const bottomBorderCase = isLastGroup && isLastCase ? "border-b-0" : ""
    const bottomBorderGroup = isLastGroup ? "border-b-0" : ""
    return (
        <tr>
            {index === 0 && (
                <>
                    <td
                        className={`py-2.5 border border-t-0 border-dividecol text-center ${bottomBorderGroup}`}
                        rowSpan={numCases}>
                        <div className="flex flex-col justify-center">
                            <div className="flex items-center justify-center mb-2">
                                {group.failed && <FontAwesomeIcon icon="fa-xmark" className="w-6 h-6 highlight-red" />}
                                {!group.failed && group.completed && (
                                    <FontAwesomeIcon icon="fa-check" className="w-6 h-6 highlight-green" />
                                )}
                                {!group.failed && !group.completed && <SVGSpinner cls="w-6 h-6" />}
                            </div>
                            {group.name}
                        </div>
                    </td>
                    <td
                        className={`py-2.5 border border-t-0 border-dividecol text-center ${bottomBorderGroup}`}
                        rowSpan={numCases}>
                        {`${group.score} / ${group.maxScore}`}
                    </td>
                </>
            )}
            <td className={`py-2.5 border border-t-0 border-dividecol text-center ${bottomBorderCase}`}>
                {testCase.index}
            </td>
            {group.scoring !== 1 && (
                <td className={`py-2.5 border border-t-0 border-dividecol ${bottomBorderCase}`} colSpan={2}>
                    <div className="flex">
                        {testCase.verdictType === 0 && <SVGSpinner cls="w-4 h-4 mr-3" />}
                        {testCase.verdictType === 1 && (
                            <FontAwesomeIcon icon="fa-xmark" className="w-4 h-4 highlight-red mr-3" />
                        )}
                        {testCase.verdictType === 2 && (
                            <FontAwesomeIcon icon="fa-check" className="w-4 h-4 highlight-yellow mr-3" />
                        )}
                        {testCase.verdictType === 3 && (
                            <FontAwesomeIcon icon="fa-check" className="w-4 h-4 highlight-green mr-3" />
                        )}
                        <span className="whitespace-nowrap">{testCase.verdictName}</span>
                    </div>
                </td>
            )}
            {group.scoring === 1 && (
                <>
                    <td className={`py-2.5 border border-t-0 border-dividecol ${bottomBorderCase}`}>
                        <div className="flex items-center">
                            <FontAwesomeIcon icon="fa-xmark" className="w-4 h-4 highlight-red mr-3" />
                            <span className="whitespace-nowrap">{testCase.verdictName}</span>
                        </div>
                    </td>
                    <td
                        className={`py-2.5 border border-t-0 border-dividecol text-center whitespace-nowrap ${bottomBorderCase}`}>
                        {testCase.score} / {testCase.maxScore}
                    </td>
                </>
            )}
            <td className={`py-2.5 border border-t-0 border-dividecol ${bottomBorderCase}`}>{testCase.time} ms</td>
            <td className={`py-2.5 border border-t-0 border-r-0 border-dividecol ${bottomBorderCase}`}>
                {testCase.memory} KiB
            </td>
        </tr>
    )
}

function TestGroup({ group, isLast }) {
    const testCases = group.testCases
    const testCasesContent = testCases.map((testCase, index) => (
        <TestCase13
            index={index}
            numCases={testCases.length}
            testCase={testCase}
            group={group}
            key={index}
            isLastGroup={isLast}
            isLastCase={index === testCases.length - 1}
        />
    ))
    return <>{testCasesContent}</>
}

function TestCase0({ testCase, onRowClicked }) {
    const { t } = useTranslation()
    return (
        <tr className="cursor-pointer hover:bg-grey-825" onClick={() => onRowClicked(testCase)}>
            <td className="py-3">
                <div className="flex items-center">
                    {testCase.verdictType === 0 && <SVGSpinner cls="w-4 h-4 mr-3" />}
                    {testCase.verdictType === 1 && (
                        <FontAwesomeIcon icon="fa-xmark" className="w-4 h-4 mr-3 highlight-red" />
                    )}
                    {testCase.verdictType === 2 && (
                        <FontAwesomeIcon icon="fa-check" className="w-4 h-4 mr-3 highlight-yellow" />
                    )}
                    {testCase.verdictType === 3 && (
                        <FontAwesomeIcon icon="fa-check" className="w-4 h-4 mr-3 highlight-green" />
                    )}
                    <span>{testCase.verdictName}</span>
                </div>
            </td>
            <td className="w-0 py-3">
                <button aria-label={t("aria_label.view")}>
                    <SVGView cls="w-6 h-6 text-grey-200 hover:text-indigo-500" />
                </button>
            </td>
            <td className="py-3">{testCase.time} ms</td>
            <td className="py-3">{testCase.memory} KiB</td>
        </tr>
    )
}

function SubmissionTable0({ status }) {
    const { t } = useTranslation()
    const [testCase, setTestCase] = useState(null)
    const [isModalOpen, setModalOpen] = useState(false)

    const handleRowClicked = (testCase) => {
        console.log(testCase.output)
        setTestCase(testCase)
        setModalOpen(true)
    }
    const testCases = status.groups?.[0].testCases
    const testCasesContent = testCases.map((testCase, index) => (
        <TestCase0 testCase={testCase} onRowClicked={handleRowClicked} key={index} />
    ))
    const outputRows = [
        ["submission_table.output", testCase?.output],
        ["submission_table.expected_output", testCase?.expectedOutput],
        ["submission_table.checker_output", testCase?.checkerOutput],
    ].map((item, index) => (
        <tr key={index}>
            <td className="whitespace-nowrap max-w-xl">{t(item[0])}</td>
            <td className="p-0 bg-codebgcol w-96">
                <CopyableCode cls="border-0 rounded-none" text={item[1]} maxHeight={"6rem"} />
            </td>
        </tr>
    ))
    const titleComponent = (
        <div className="py-4 px-5 flex justify-center items-center space-x-2 border-b border-bordefcol">
            <div className="flex items-center">
                {testCase?.verdictType === 0 && <SVGSpinner cls="w-5 h-5 mr-3" />}
                {testCase?.verdictType === 1 && (
                    <FontAwesomeIcon icon="fa-xmark" className="w-5 h-5 highlight-red mr-3" />
                )}
                {testCase?.verdictType === 2 && (
                    <FontAwesomeIcon icon="fa-check" className="w-5 h-5 highlight-yellow mr-3" />
                )}
                {testCase?.verdictType === 3 && (
                    <FontAwesomeIcon icon="fa-check" className="w-5 h-5 highlight-green mr-3" />
                )}
                <div className="space-x-2">
                    <span>{t("submission_table.test_case")}</span>
                    <span>#{testCase?.index}</span>
                </div>
            </div>
        </div>
    )
    return (
        <>
            <Modal isOpen={isModalOpen} onClose={() => setModalOpen(false)}>
                <RoundedTable cls="w-full sm:w-[30rem]" titleComponent={titleComponent}>
                    <tbody>{outputRows}</tbody>
                </RoundedTable>
            </Modal>
            <RoundedTable>
                <thead>
                    <tr>
                        <th colSpan={2}>{t("submission_table.verdict")}</th>
                        <th>{t("submission_table.time")}</th>
                        <th>{t("submission_table.memory")}</th>
                    </tr>
                </thead>
                <tbody>{testCasesContent}</tbody>
            </RoundedTable>
        </>
    )
}

function SubmissionTable13({ status }) {
    const { t } = useTranslation()
    const groups = status.groups
    const groupsContent = groups.map((group, index) => (
        <TestGroup group={group} isLast={index === groups.length - 1} key={index} />
    ))
    return (
        <RoundedTable>
            <thead>
                <tr>
                    <th>{t("submission_table.subtask")}</th>
                    <th>{t("submission_table.total")}</th>
                    <th>{t("submission_table.test")}</th>
                    <th colSpan={2}>{t("submission_table.verdict")}</th>
                    <th>{t("submission_table.time")}</th>
                    <th>{t("submission_table.memory")}</th>
                </tr>
            </thead>
            <tbody>{groupsContent}</tbody>
        </RoundedTable>
    )
}
function SubmissionTable({ status }) {
    if (status.feedbackType === 0) {
        return <SubmissionTable0 status={status} />
    }
    if ([1, 3].includes(status.feedbackType)) {
        return <SubmissionTable13 status={status} />
    }
}

export default SubmissionTable
