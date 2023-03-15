import React from "react";
import {makeStyles} from "@material-ui/core";
import timeLogo from "../Assets/Logo/timeLoading.gif";
// import { useNavigate } from "react-router-dom";

const useStyles = makeStyles({
    root: {
       border: "5px solid #5e5858",
       borderRadius: "8px",
       width: "100%vh",
       height: "98vh",
    },
    menuContainer: {
       listStyleType: "none",
       color: "white",
       margin: "0px",
       padding: "0px",
       width: "250px",
       height: "98vh",
       backgroundColor: "#1c1b1b",
   },
   logoContainer: {
       height: "150px"
   },
   listStyle: {
       display: "block",
       padding: "0px 0px 0px 0px",
   },
   logo: {
    width: "50px",
    height: "50px",
    margin: "35px 0px 0px 20px"
   },
   menuField: {
    display: "block",
    color:"gray",
    width: "250px",
    height: "50px",
    alignItems: "center",
    textDecoration: "none",
    "&:hover":{
        color: "white",
        cursor: "pointer",
        backgroundColor: "#575353",
    }
   },
   menufieldValue: {
    marginLeft: "65px",
    alignItems: "center",
    textAlign: "center",
    justifyContent: "center",
    paddingTop: "15px"
   }
})

const DashboardLayout = () => {
    const style = useStyles();
    // const navigate = useNavigate();
    return (
        <div className={style.root}>
          <div className={style.menuContainer}>
              <div className={style.logoContainer}>
                <img className={style.logo} src={timeLogo} alt="timeLogo"/>
              </div>
                <li className={style.listStyle}
                    // onClick={() => navigate('/dashboard')}
                >
                  <span className={style.menuField}><span className={style.menufieldValue}>Dashboard</span></span>
                </li>
                <li className={style.listStyle} 
                    // onClick={() => navigate('/task')}
                >
                  <span className={style.menuField}><span className={style.menufieldValue}>Task</span></span>
                </li>
                <li className={style.listStyle} 
                    // onClick={() => navigate('/project')}
                >
                  <span className={style.menuField}><span className={style.menufieldValue}>Project</span></span>
                </li>
                <li className={style.listStyle} 
                    // onClick={() => navigate('/notification')}
                >
                  <span className={style.menuField}><span className={style.menufieldValue}>Notifications</span></span>
                </li>
                <li className={style.listStyle} 
                    // onClick={() => navigate('/settings')}
                >
                  <span className={style.menuField}><span className={style.menufieldValue}>Settings</span></span>
                </li> 
          </div>
        </div>
    )
};

export default DashboardLayout;