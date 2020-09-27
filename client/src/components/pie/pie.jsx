import React from "react";
import { Doughnut } from "react-chartjs-2";

const data = {
  labels: ["In Progress", "Completed", "Paused"],
  datasets: [],
};

const Bie = React.memo(({ dataSet }) => {
  console.log(dataSet);
  data.datasets[0] = dataSet;
  return (
    <>
      <div
        style={{
          position: "fixed",
          width: "700px",
          left: "43rem",
          bottom: "15rem",
        }}
      >
        <Doughnut data={data} />
      </div>
    </>
  );
});

export default Bie;
