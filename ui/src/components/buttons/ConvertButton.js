import React from "react";
import { useState } from "react";
import { instance } from "../../utils/Utils";

const ConvertButton = () => {
  
  const [isConverting, setIsConverting] = useState(false);
  const [errorExist, setErrorExist] = useState(false);

  const clickHandler = async () => {
    setIsConverting(true);
    const res = await instance.get("/convert");
    const status_code = await res.status;
    if (status_code === 200) {
      setIsConverting(false);
    } else {
      setErrorExist(true);
      setIsConverting(false);
    }
  };

  return (
    <React.Fragment>
      <button
        className="block w-1/2 m-auto mb-2 font-mono font-bold bg-blue-400 border-2 border-black"
        onClick={clickHandler}
        disabled={isConverting}
      >
        convert
      </button>
      {errorExist ? (
        <div className="-mt-2 font-mono font-bold text-center text-red-400">
          Source file not found!
        </div>
      ) : null}
    </React.Fragment>
  );
};

export default ConvertButton;
