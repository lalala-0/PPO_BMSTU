import React from "react";
import { Link } from "react-router-dom";
import "./Header.css";

const Header: React.FC = () => {
    return (
        <header className="header">
            <nav>
                <Link to="/ratings" className="nav-link">
                    Список рейтингов
                </Link>
            </nav>
        </header>
    );
};

export default Header;
