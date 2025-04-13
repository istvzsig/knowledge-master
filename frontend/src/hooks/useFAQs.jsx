import { useState, useEffect } from "react";
import useSessionStorage from "../hooks/useSessionStorage";

import * as config from "../config";

function useFAQs(path = "/faqs") {
  const sessionStorage = useSessionStorage();
  const [faqs, setFaqs] = useState(sessionStorage.load(path) || []);
  const [error, setError] = useState(undefined);

  let loading = false;
  let currentIndex = 0;

  async function fetchFAQs() {
    try {
      const response = await fetch(path, config.headers.faqs);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return await response.json();
    } catch (err) {
      console.error("Error fetching FAQs:", err);
      setError(err);
    }
  }

  async function loadFAQs() {
    if (faqs.length === 0) {
      loading = true;
      const faqs = await fetchFAQs();
      setFaqs(faqs);
      sessionStorage.save(path, faqs);
      loading = false;
    }
  }

  useEffect(() => {
    loadFAQs();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [path]);

  return { faq: faqs[currentIndex], faqs, error, loading };
}

export default useFAQs;
