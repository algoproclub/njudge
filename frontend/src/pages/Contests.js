import React from 'react';
import ProfileSideBar from '../components/concrete/other/ProfileSidebar'
import ContestList from '../components/concrete/other/ContestList'
import '../index.css';
import checkData from "../util/CheckData";

function Contests({data}) {
    if (!checkData(data)) {
        return
    }
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar profileData={data.profileData}/>
                </div>
                <div className="w-full px-4 lg:pl-3">
                    <ContestList contestData={[
                        ["Online oktató programozóverseny #3", "2023-12-23, 14:00", true],
                        ["Online oktató programozóverseny #2", "2023-08-23, 14:00", false],
                        ["Online oktató programozóverseny #1", "2023-04-23, 14:00", false]]}/>
                </div>
            </div>
        </div>
    )
}

export default Contests;