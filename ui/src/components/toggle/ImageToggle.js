import { useContext } from "react";
import Context from "../../store/context";

const ImageToggle = () => {
  const { imageDisplayToggle, setImageDisplayToggle } = useContext(Context);

  return (
    <label className="flex items-center justify-between">
      > show image
      <input
        id="image-toggle"
        type="checkbox"
        className="w-5 h-5 ml-2"
        defaultChecked={imageDisplayToggle}
        onClick={() => {
          setImageDisplayToggle(!imageDisplayToggle);
        }}
      />
    </label>
  );
};

export default ImageToggle;
