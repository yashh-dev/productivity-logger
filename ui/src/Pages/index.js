import React from "react";
import { connect } from "react-redux";
import DashboardLayout from "../Layout/dashboardLayout";
import store from "../Store/storeIndex";
import Dashboard from "./Dashboard/dashboard";
import Notifications from "./Notification/notification";


const Home = ({menu}) => {
    const [value, setValue] = React.useState(0);
    React.useEffect(()=> {
       console.log(menu)
    },[menu]);  
    
    const route = () => {
       switch (value){
        case 0:
            return <Dashboard/>;
        case 1:
            return <Notifications/>;
        default:
            return <Dashboard/>;
       }
    }

    return(
        <DashboardLayout>
          <Dashboard/>
        </DashboardLayout>
    )
};

const mapStateToProps = (state) => ({
    menu: state
})

export default connect(mapStateToProps)(Home);