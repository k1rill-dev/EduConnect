import React from 'react';
import {Navigate} from "react-router-dom";
// import api from "../../api";
//
// const logoutUser = async () => {
//     await api.get('/api/logout', {
//         withCredentials: true,
//     })
// }

const logoutUser = async () => {

}

const LogoutForm = () => {
    logoutUser();
    localStorage.clear();
    return (
        <div><Navigate to="/"/></div>
    )
};

export default LogoutForm;