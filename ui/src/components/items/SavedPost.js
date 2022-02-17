import { forwardRef, useContext } from "react";
import ReactMarkdown from "react-markdown";
import BookmarkButton from "../buttons/BookmarkButton";
import Context from "../../store/context";

const SavedPost = forwardRef((props, ref) => {
  const {
    title,
    body,
    timestamp,
    link,
    url,
    saved,
    subreddit,
    author,
    name,
    concise,
  } = props;
  const { imageDisplayToggle } = useContext(Context);

  const parsedDate = new Date(timestamp * 1000).toLocaleString();
  const domains = ["i.redd.it", "storage.googleapis.com", "reddit.com/gallery"];

  let parsedURL = "";

  try {
    parsedURL = new URL(url);
  } catch (err) {
    return false;
  }

  let filtered_url = url;
  if (!domains.includes(parsedURL.hostname)) {
    filtered_url = "/empty";
  }

  return (
    <div>
      <div
        ref={ref}
        className="flex-1 p-3 font-mono bg-green-200 border-2 border-black border-dashed hover:border-solid hover:border-3"
      >
        <div className="grid grid-cols-3">
          <div className="h-4 text-xs leading-3 col-span-2">
            <p className="mb-[0.075rem] font-bold">r/{subreddit}</p>
            <h1 className="tracking-tight">by u/{author}</h1>
          </div>
          {concise ? (
            <div className="h-7"></div>
          ) : (
            <BookmarkButton saved={saved} name={name} />
          )}
        </div>

        <a href={link} target="_blank" rel="noreferrer">
          <div>
            <h2 className="mr-2 font-bold break-normal leading-5">{title}</h2>
            <p className="text-xs tracking-tighter">{parsedDate}</p>
            <div className="p-2 overflow-hidden text-sm text-center">
              <ReactMarkdown>
                {body.length < 499 ? body : body + "..."}
              </ReactMarkdown>
            </div>
            { 
              /*
              Nampilin gambar kalo concise === false 
              Sama kalo imageDisplayToggle === true 
              */ 
            }
            { concise ? null : imageDisplayToggle ? (
              <div className="relative w-full h-full">
                <img src={filtered_url} loading="lazy" alt="" />
              </div>
            ) : null}
          </div>
        </a>
      </div>
    </div>
  );
});

export default SavedPost;
