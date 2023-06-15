import React from "react";
import { makeStyles } from "@material-ui/core";
import Logo from "../Assets/Logo/Logo.png";
// import store from "../Store/storeIndex";
import { connect } from "react-redux";
import { HOME, PROJECT, TASK, NOTIFICATION, SETTINGS } from "../Routes/common";
// import { setMenu } from "../Store/Action/DashBoard/dashboardAction";
import { useNavigate } from "react-router-dom";

const useStyles = makeStyles({
  base: {
    border: "5px solid gray",
    borderRadius: "8px",
    height: "98vh",
  },
  root: {
    height: "98vh",
    display: "flex",
  },
  menuContainer: {
    listStyleType: "none",
    color: "white",
    margin: "0px",
    padding: "0px",
    width: "250px",
    height: "98vh",
    backgroundColor: "#1c1b1b",
    display: "block",
  },
  logoContainer: {
    height: "150px",
  },
  listStyle: {
    display: "block",
    padding: "0px 0px 0px 0px",
  },
  logo: {
    width: "35px",
    height: "35px",
    margin: "35px 0px 0px 20px",
  },
  menuField: {
    display: "flex",
    color: "gray",
    width: "250px",
    height: "50px",
    alignItems: "center",
    textDecoration: "none",
    "&:hover": {
      color: "white",
      cursor: "pointer",
      backgroundColor: "#575353",
    },
  },
  menufieldValue: {
    marginLeft: "65px",
    alignItems: "center",
    textAlign: "center",
    justifyContent: "center",
    paddingTop: "15px",
  },
});

const DashboardLayout = ({ children }) => {
  const style = useStyles();
  const navigate = useNavigate();
  // const [menu, setMenu] = React.useState(0);
  // console.log("Menu...",menu);
  //  React.useEffect(()=> {
  //  store.dispatch(setMenu(menu))
  //  },[menu])
  return (
    <div className={style.base}>
      <div className={style.root}>
        <div className={style.menuContainer}>
          <div className={style.logoContainer}>
            <img className={style.logo} src={Logo} alt="timeLogo" />
          </div>
          <span className={style.listStyle} onClick={() => navigate(HOME)}>
            <span className={style.menuField}>
              <span className={style.menufieldValue}>Dashboard</span>
            </span>
          </span>
          <span className={style.listStyle} onClick={() => navigate(TASK)}>
            <span className={style.menuField}>
              <span className={style.menufieldValue}>Task</span>
            </span>
          </span>
          <span className={style.listStyle} onClick={() => navigate(PROJECT)}>
            <span className={style.menuField}>
              <span className={style.menufieldValue}>Project</span>
            </span>
          </span>
          <span
            className={style.listStyle}
            onClick={() => navigate(NOTIFICATION)}
          >
            <span className={style.menuField}>
              <span className={style.menufieldValue}>Notifications</span>
            </span>
          </span>
          <span className={style.listStyle} onClick={() => navigate(SETTINGS)}>
            <span className={style.menuField}>
              <span className={style.menufieldValue}>Settings</span>
            </span>
          </span>
        </div>
        <div>{children}</div>
      </div>
    </div>
  );
};

const mapStateToProps = (state) => ({
  menu: state,
});

export default connect(mapStateToProps)(DashboardLayout);
