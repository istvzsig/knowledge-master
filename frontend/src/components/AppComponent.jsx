import useFAQs from "../hooks/useFAQs";

import FAQComponent from "./faq/FAQComponent";

export default function App() {
  const { faq, loading } = useFAQs();

  return (
    <div>
      {loading ? (
        <h1>Loading...</h1>
      ) : (
        <div>
          {faq && <FAQComponent faq={faq} />}
          <button onClick={null}>ASD</button>
        </div>
      )}
    </div>
  );
}
