import React from "react";
import { makeStyles } from "@material-ui/core";
import { IconButton } from "@material-ui/core";
import PlayArrowIcon from "@mui/icons-material/PlayArrow";
import PauseIcon from "@mui/icons-material/Pause";
import SaveAltRoundedIcon from "@mui/icons-material/SaveAltRounded";

const useStyles = makeStyles({
  root: {
    display: "flex",
  },
  pauseResumeBtn: {
    margin: "30px 3px 0px 5px",
    display: "flex",
  },
  stopwatchTime: {
    fontSize: "30px",
    width: "100px",
    height: "20px",
    marginTop: "40px",
    marginRight: "100px",
    marginLeft: "100px",
  },
  button: {
    margin: "20px",
    border: "none",
    padding: "10px 30px",
    cursor: "pointer",
    color: "orange",
    display: "flex",
    justifyContent: "center",
  },
});

const Stopwatch = ({ title }) => {
  const style = useStyles();
  const [time, setTime] = React.useState(0);
  const [isRunning, setIsRunning] = React.useState(false);
  const [record, setRecord] = React.useState({
    endTime: "",
    time: "",
    title: "",
  });

  React.useEffect(() => {
    let interval;
    if (isRunning) {
      interval = setInterval(() => setTime(time + 1), 10);
    }
    return () => clearInterval(interval);
  }, [isRunning, time]);

  const hour = Math.floor(time / 360000);
  const minute = Math.floor((time % 360000) / 6000);
  const seconds = Math.floor((time % 6000) / 100);
  const milliseconds = time % 100;

  const startAndStop = () => {
    setIsRunning(!isRunning);
  };

  const Save = () => {
    const presentTime = new Date();
    setRecord({
      time: time,
      title: title,
      endTime: presentTime.getTime(),
    });
    console.log(time, title, presentTime.getTime());
    setTime(0);
  };

  return (
    <div className={style.root}>
      <p className={style.stopwatchTime}>
        {hour.toString().padStart(2, "0")}:{minute.toString().padStart(2, "0")}:
        {seconds.toString().padStart(2, "0")}:
        {milliseconds.toString().padStart(2, "0")}
      </p>
      <div className={style.pauseResumeBtn}>
        <IconButton onClick={startAndStop}>
          {isRunning ? (
            <PauseIcon color="primary" />
          ) : (
            <PlayArrowIcon color="primary" />
          )}
        </IconButton>
        <IconButton onClick={Save} color="secondary" size="100px">
          <SaveAltRoundedIcon />
        </IconButton>
      </div>
    </div>
  );
};

export default Stopwatch;
