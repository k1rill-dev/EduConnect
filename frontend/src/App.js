import {BrowserRouter, Route, Routes} from "react-router-dom";
import './App.css';
import LogoutForm from "./components/forms/LogoutForm";
import RegistrationForm from "./components/forms/RegistrationForm";
import LoginForm from "./components/forms/LoginForm";
import Header from "./components/header/Header";
import Main from "./components/pages/Main";
import Footer from "./components/footer/Footer";
import CoursesPage from "./components/pages/CoursesPage";
import CoursePage from "./components/pages/CoursePage";
import CoursePageForStudent from "./components/pages/CoursePageForStudent";
import ProfilePage from "./components/pages/ProfilePage";
import CreateJobPage from "./components/pages/CreateJobPage";
import CreateCoursePage from "./components/pages/CreateCoursePage";
import EditProfilePage from "./components/pages/EditProfilePage";
import JobsPage from "./components/pages/JobsPage";
import JobDetailPage from "./components/pages/JobDetailPage";

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
                <Route path="/courses" element={(
                    <div className="flex flex-col min-h-screen">
                        <Header/>

                        <div className="flex-grow">
                            <CoursesPage/>
                        </div>

                        <Footer/>
                    </div>
                )}>
                </Route>
                <Route path="/courses/:courseId" element={(
                    <div className="flex flex-col min-h-screen">
                        <Header/>

                        <div className="flex-grow">
                            <CoursePage />
                        </div>

                        <Footer/>
                    </div>
                )}>
                </Route>
                <Route path="/courseStudent/:courseId" element={(
                    <div className="flex flex-col min-h-screen">
                        <Header/>

                        <div className="flex-grow">
                            <CoursePageForStudent/>
                        </div>

                        <Footer/>
                    </div>
                )}>
                </Route>
                <Route path="/profile" element={(
                    <div className="flex flex-col min-h-screen">
                        <Header/>

                        <div className="flex-grow">
                            <ProfilePage/>
                        </div>

                        <Footer/>
                    </div>
                )}>
                </Route>
                <Route path="/create-job" element={(
                    <div className="flex flex-col min-h-screen">
                        <Header/>

                        <div className="flex-grow">
                            <CreateJobPage/>
                        </div>

                        <Footer/>
                    </div>
                )}>
                </Route>
                <Route path="/create-course" element={(
                    <div className="flex flex-col min-h-screen">
                        <Header/>

                        <div className="flex-grow">
                            <CreateCoursePage/>
                        </div>

                        <Footer/>
                    </div>
                )}>
                </Route>
                <Route path="/settings" element={(
                    <div className="flex flex-col min-h-screen">
                        <Header/>

                        <div className="flex-grow">
                            <EditProfilePage/>
                        </div>

                        <Footer/>
                    </div>
                )}>
                </Route>
                <Route path="/jobs" element={(
                    <div className="flex flex-col min-h-screen">
                        <Header/>

                        <div className="flex-grow">
                            <JobsPage/>
                        </div>

                        <Footer/>
                    </div>
                )}>
                </Route>
                <Route path="/jobs/:jobId" element={(
                    <div className="flex flex-col min-h-screen">
                        <Header/>

                        <div className="flex-grow">
                            <JobDetailPage/>
                        </div>

                        <Footer/>
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
