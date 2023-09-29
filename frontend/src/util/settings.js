import { apiRoute } from "../config/RouteConfig";

export async function saveSettings(user, showUnsolved, hideSolved) {
    const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            showUnsolved: showUnsolved,
            hideSolved: hideSolved,
        }),
        credentials: 'include',
    };
    console.log(
        apiRoute(`/user/profile/${encodeURIComponent(user)}/settings/other/`),
    );
    const response = await fetch(
        apiRoute(`/user/profile/${encodeURIComponent(user)}/settings/other/`),
        requestOptions,
    );
    const data = await response.json();
    return { ...data, success: response.ok };
}

export async function changePassword(user, oldPw, newPw, newPwConfirm) {
    const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            oldPw: oldPw,
            newPw: newPw,
            newPwConfirm: newPwConfirm,
        }),
        credentials: 'include',
    };
    const response = await fetch(
        apiRoute(
            `/user/profile/${encodeURIComponent(
                user,
            )}/settings/change_password/`,
        ),
        requestOptions,
    );
    const data = await response.json();
    return { ...data, success: response.ok };
}
