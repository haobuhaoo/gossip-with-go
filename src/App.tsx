import LoginPage from "./pages/login";
import HomePage from "./pages/home";
import TopicPage from "./pages/topic";
import PostPage from "./pages/post";
import { BrowserRouter, Route, Routes } from "react-router-dom";

const App: React.FC = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<LoginPage />} />
                <Route path="/home" element={<HomePage />} />
                <Route path="/home/:id" element={<TopicPage />} />
                <Route path="/home/:id/:id" element={<PostPage />} />
            </Routes>
        </BrowserRouter>
    );
};

export default App;
