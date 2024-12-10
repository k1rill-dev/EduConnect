import {BrowserRouter, Route, Routes} from "react-router-dom";
import './App.css';
import LogoutForm from "./components/forms/LogoutForm";
import RegistrationForm from "./components/forms/RegistrationForm";
import LoginForm from "./components/forms/LoginForm";
import Header from "./components/header/Header";
import Main from "./components/pages/Main";
import Footer from "./components/footer/Footer";

function App() {
 return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={(
                    <div>
                        <Header></Header>
                        <Main></Main>
                        <Footer></Footer>
                    </div>
                )}>
                </Route>
                <Route path="/logout" element={(
                    <div>
                        <LogoutForm></LogoutForm>
                    </div>
                )}>
                </Route>
                <Route path="/register" element={(
                    <div>
                        <RegistrationForm></RegistrationForm>
                    </div>
                )}>
                </Route>
                <Route path="/login" element={(
                    <div>
                        <LoginForm></LoginForm>
                    </div>
                )}>
                </Route>
            </Routes>
        </BrowserRouter>
    );
}

export default App;
