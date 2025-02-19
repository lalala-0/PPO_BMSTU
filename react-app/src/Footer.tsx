import React from "react";
import { Link } from "react-router-dom";
import { isAuthenticated } from "./ui/controllers/autchControllers/isAutchenticated";
import { logout } from "./ui/controllers/autchControllers/logout";

const Footer: React.FC = () => {
    return (
        <footer className="footer">
            {isAuthenticated() ? (
                <>
                    <Link to="/judgeDashboard" className="footer-link">
                        Панель управления
                    </Link>
                    <Link to="/profile" className="footer-link">
                        Профиль
                    </Link>
                    <button onClick={logout} className="footer-button">
                        Выйти
                    </button>
                </>
            ) : (
                <Link to="/login" className="footer-link">
                    Войти
                </Link>
            )}
            <p>&copy; 2025, Все права защищены.</p>
        </footer>
    );
};

export default Footer;
