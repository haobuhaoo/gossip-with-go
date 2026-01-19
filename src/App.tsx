import { BrowserRouter, Route, Routes } from "react-router-dom";

import AuthProvider from "./context/authprovider";
import ProtectedRoute from "./context/protectedroute";

import HomePage from "./pages/home";
import LoginPage from "./pages/login";
import PostPage from "./pages/post";
import TopicPage from "./pages/topic";

import "./App.css";

const App: React.FC = () => {
    return (
        <BrowserRouter>
            <AuthProvider>
                <Routes>
                    <Route path="/" element={<LoginPage />} />
                    <Route element={<ProtectedRoute />}>
                        <Route path="/home" element={<HomePage />} />
                        <Route path="/home/:topicId" element={<TopicPage />} />
                        <Route path="/home/:topicId/:postId" element={<PostPage />} />
                    </Route>
                </Routes>
            </AuthProvider>
        </BrowserRouter>
    );
};

export default App;
