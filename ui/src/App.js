import './App.css';
import {BrowserRouter, Routes, Route} from 'react-router-dom';
import DashboardLayout from './Layout/dashboardLayout';
import Settings from './Pages/Settings/settings';
import Notifications from './Pages/Notification/notification';
import Projects from './Pages/Project/projects';
import Task from './Pages/Tasks/task';
import Dashboard from './Pages/Dashboard/dashboard';

function App() {
  return (
   <BrowserRouter>
     <Routes>
       <Route path='/' element={<Dashboard/>}/>
       <Route path='/task' element={<Task/>}/>
       <Route path='/project' element={<Projects/>}/>
       <Route path='/notification' element={<Notifications/>}/>
       <Route path='/settings' element={<Settings/>}/>
     </Routes>
   </BrowserRouter>
  );
}

export default App;
