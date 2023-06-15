import React from "react";
import DashboardLayout from "../../Layout/dashboardLayout";
import { makeStyles, TextField } from "@material-ui/core";
import Stopwatch from "../../Components/Stopwatch";

const useStyles = makeStyles({
  titleTimeBar: {
    //   "&:css-1a08mon-TimerFormContent": {
    //     display: "flex",
    //     flexDirection: "row",
    //     alignItem: "center",
    //     flexBasis: "84px",
    // }
    width: "100vh",
    // backgroundColor: "gray",
    marginTop: "1px",
    height: "100px",
    display: "flex",
  },
  textField: {
    width: "100vh",
    height: "80px",
    margin: "45px 280px 0px 15px",
    fontSize: "50px",
    border: "0.5px solid #aaa5b0",
    borderRadius: "5px",
  },
  root: {
    marginTop: 0,
    display: "flex",
  },
  titleField: {
    width: "50vh",
  },
  stopwatch: {
    marginLeft: "60vh",
  },
});

const Dashboard = () => {
  const style = useStyles();
  const [title, setTitle] = React.useState("");

  const onChange = (e) => {
    setTitle(e.target.value);
  };

  return (
    <DashboardLayout>
      <div className={style.root}>
        <div className={style.titleTimeBar}>
          <div className={style.titleField}>
            <TextField
              className={style.textField}
              fullWidth
              name="title"
              value={title}
              placeholder=" Enter your work title..."
              InputProps={{ disableUnderline: true }}
              onChange={onChange}
            />
          </div>
          <div className={style.stopwatch}>
            <Stopwatch title={title} />
          </div>
        </div>
      </div>
    </DashboardLayout>
  );
};

export default Dashboard;
