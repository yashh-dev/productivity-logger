import './App.css';
import {BrowserRouter, Routes, Route} from 'react-router-dom';
import DashboardLayout from './Layout/dashboardLayout';
import Settings from './Pages/settings';
import Notifications from './Pages/notification';
import Projects from './Pages/projects';
import Task from './Pages/task';
import Dashboard from './Pages/dashboard';

function App() {
  return (
   <BrowserRouter>
     <Routes>
       <Route path='/' element={<DashboardLayout/>}/>
       <Route path='/dashboard' element={<Dashboard/>}/>
       <Route path='/task' element={<Task/>}/>
       <Route path='/project' element={<Projects/>}/>
       <Route path='/notification' element={<Notifications/>}/>
       <Route path='/settings' element={<Settings/>}/>
     </Routes>
   </BrowserRouter>
  );
}

export default App;
