import { useContext, useState } from "react";
import { Link, Navigate, useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import RoundedFrame from "../../components/container/RoundedFrame";
import TextBox from "../../components/input/TextBox";
import SVGTitleComponent from "../../components/svg/SVGTitleComponent";
import { login } from "../../util/auth";
import { routeMap } from "../../config/RouteConfig";
import UserContext from "../../contexts/user/UserContext";

function LoginFrame() {
    const { t } = useTranslation();
    const navigate = useNavigate();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const titleComponent = (
        <SVGTitleComponent
            svg={<FontAwesomeIcon icon="fa-sign-in" className="w-5 h-5 mr-3" />}
            title={t("login.login")}
        />
    );
    const handleLogin = (event) => {
        event.preventDefault();
        login(username, password).then((resp) => {
            if (resp.success) {
                window.flash("flash.successful_login", "success");
                navigate(routeMap.home);
            } else {
                window.flash("flash.unsuccessful_login", "failure");
            }
        });
    };
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-10 py-8">
                <form method="POST" className="mb-4">
                    <div className="mb-4">
                        <TextBox
                            id="username"
                            label={t("login.username")}
                            initText={username}
                            onChange={setUsername}
                        />
                    </div>
                    <div className="mb-2">
                        <TextBox
                            id="password"
                            label={t("login.password")}
                            initText={password}
                            type="password"
                            onChange={setPassword}
                        />
                    </div>
                    <div className="mb-6">
                        <Link to={routeMap.forgotten_password} className="link">
                            {t("login.forgotten_password")}
                        </Link>
                    </div>
                    <div className="flex justify-center mb-1">
                        <button
                            type="submit"
                            className="btn-indigo padding-btn-default mr-2 w-1/2"
                            onClick={handleLogin}>
                            {t("login.login")}
                        </button>
                        <button className="btn-gray padding-btn-default flex items-center justify-center w-1/2">
                            Google
                        </button>
                    </div>
                </form>
            </div>
        </RoundedFrame>
    );
}

function Login() {
    const { userData, isLoggedIn } = useContext(UserContext);
    if (isLoggedIn) {
        return (
            <Navigate
                to={routeMap.profile.replace(
                    ":user",
                    encodeURIComponent(userData.username),
                )}
            />
        );
    }
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full sm:max-w-md">
                <div className="w-full px-4">
                    <LoginFrame />
                </div>
            </div>
        </div>
    );
}

export default Login;
