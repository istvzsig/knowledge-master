import { useState, useEffect } from "react";
import useSessionStorage from "../hooks/useSessionStorage";

import * as config from "../config";

function useFAQs(path = "/faqs") {
  const sessionStorage = useSessionStorage();
  const [faqs, setFaqs] = useState(sessionStorage.load(path) || []);
  const [error, setError] = useState(undefined);

  let loading = false;
  let currentFAQIndex = 0;

  async function fetchFAQs() {
    try {
      const response = await fetch(path, config.headers.faqs);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      data.forEach(
        (d) => (d.createdAt = new Date(d.createdAt * 1000).toLocaleString())
      );

      return data;
    } catch (err) {
      console.error("Error fetching FAQs:", err);
      setError(err);
    }
  }

  function getNextFAQ() {
    if (faqs.length > 0) {
      currentFAQIndex = Math.floor((currentFAQIndex + 1) % faqs.length);
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

  const nextFAQ = faqs[currentFAQIndex];

  return { faqs, error, loading, nextFAQ, getNextFAQ };
}

export default useFAQs;
