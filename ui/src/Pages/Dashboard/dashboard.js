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
    display: "flex"
  },
  textField: {
     width: "100vh",
     height: "80px",
     margin: "45px 280px 0px 15px",
     fontSize: "50px"
  }
});

const Dashboard = () => {
  const style = useStyles();
  const [title, setTitle] = React.useState('');

  const onChange = (e) => {
    setTitle(e.target.value);
  };

    return (
        <DashboardLayout>
          <div style={{display: "flex"}}>
            <div className={style.titleTimeBar}>
              <div style={{width: "50vh"}}>
                 <TextField
                    className={style.textField}
                    fullWidth
                    name="title"
                    value={title}
                    placeholder="Enter your work title..."
                    InputProps={{ disableUnderline: true}}
                    onChange={onChange}
                 />
              </div>
              <div style={{marginLeft: "60vh"}}>
                  <Stopwatch title={title}/>
              </div>
            </div>
          </div>
        </DashboardLayout>
    )
};

export default Dashboard;