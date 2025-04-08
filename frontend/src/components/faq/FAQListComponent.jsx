import FAQComponent from "./FAQComponent";

export default function FAQListComponent({ faqs }) {
  return (
    <div>
      {faqs.map((faq) => (
        <FAQComponent key={faq.id} faq={faq} />
      ))}
    </div>
  );
}
