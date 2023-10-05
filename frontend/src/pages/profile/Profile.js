import { useContext } from "react"
import { Outlet, useParams } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { routeMap } from "../../config/RouteConfig"
import UserContext from "../../contexts/user/UserContext"
import TabFrame from "../../components/container/TabFrame"

function Profile() {
    console.log("nigga")
    let routeLabels = [
        "profile.profile",
        "profile.submissions",
        "profile.settings",
    ]
    let routePatterns = [
        routeMap.profile,
        routeMap.profileSubmissions,
        routeMap.profileSettings,
    ]
    const { t } = useTranslation()
    const { user } = useParams()
    const { userData, isLoggedIn } = useContext(UserContext)
    const routes = routePatterns.map((item) => item.replace(":user", user))

    if (!isLoggedIn || userData.username !== user) {
        routeLabels.pop()
        routePatterns.pop()
    }
    return (
        <div className="flex justify-center">
            <div className="w-full max-w-7xl">
                <div className="w-full px-4">
                    <TabFrame
                        routes={routes}
                        routePatterns={routePatterns}
                        routeLabels={routeLabels.map(t)}>
                        <div className="relative w-full">
                            <Outlet />
                        </div>
                    </TabFrame>
                </div>
            </div>
        </div>
    )
}

export default Profile
