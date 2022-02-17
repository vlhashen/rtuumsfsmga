import React from "react";
import { useState } from "react";
import { instance } from "../../utils/Utils";

const SyncButton = () => {

  const [isSyncing, setIsSyncing] = useState(false);

  const clickHandler = async () => {
    const res = await instance.get("/sync");
    const status_code = await res.status;
    if (status_code === 200) {
      setIsSyncing(false);
    } else {
      setIsSyncing(false);
    }
  };

  return (
    <React.Fragment>
      <button
        className="block w-1/2 m-auto mb-2 font-mono font-bold bg-blue-400 border-2 border-black"
        onClick={clickHandler}
        disabled={isSyncing}
      >
        sync
      </button>
    </React.Fragment>
  );
};

export default SyncButton;
