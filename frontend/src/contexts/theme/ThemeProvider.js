import { useEffect, useState } from "react";

import ThemeContext from "./ThemeContext";

function ThemeProvider({ children }) {
    const [theme, setTheme] = useState(
        localStorage.getItem("theme") || "light",
    );

    useEffect(() => {
        changeTheme(theme);
    }, []);

    const changeTheme = (newTheme) => {
        if (!["light", "dark"].includes(theme)) {
            return;
        }
        setTheme(newTheme);

        const doc = document.documentElement;
        if (newTheme === "light" && doc.classList.contains("dark")) {
            doc.classList.remove("dark");
        }
        if (newTheme === "dark" && !doc.classList.contains("dark")) {
            doc.classList.add("dark");
        }
        localStorage.setItem("theme", newTheme);
    };
    return (
        <ThemeContext.Provider value={{ theme, changeTheme }}>
            {children}
        </ThemeContext.Provider>
    );
}

export default ThemeProvider;
