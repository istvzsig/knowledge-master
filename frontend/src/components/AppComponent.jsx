import useFAQs from "../hooks/useFAQs";
import FAQComponent from "./faq/FAQComponent";

export default function App() {
  const { loading, nextFAQ } = useFAQs("faqs");
  return (
    <div>{loading ? <h1>Loading...</h1> : <FAQComponent faq={nextFAQ} />}</div>
  );
}
