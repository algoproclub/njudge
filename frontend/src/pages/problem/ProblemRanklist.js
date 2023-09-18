import Ranklist from "../../components/concrete/other/Ranklist";
import SVGTitleComponent from '../../svg/SVGTitleComponent';
import {SVGResults} from "../../svg/SVGs";
import Pagination from "../../components/util/Pagination";
import {matchPath, useLocation, useOutletContext} from "react-router-dom";
import checkData from "../../util/CheckData";

function ProblemRanklist() {
    const data = useOutletContext()
    const location = useLocation()
    if (!data || !matchPath(data.route, location.pathname)) {
        return
    }
    const titleComponent = <SVGTitleComponent svg={<SVGResults/>} title="Eredmények"/>
    return (
        <div>
            <div className="mb-2">
                <Ranklist ranklist={data.ranklist} titleComponent={titleComponent}/>
            </div>
            <Pagination paginationData={data.paginationData}/>
        </div>
    )
}

export default ProblemRanklist;