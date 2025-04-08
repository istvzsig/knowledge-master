import { useState, useEffect } from "react";
import useSessionStorage from "../hooks/useSessionStorage";

function useFAQs(url) {
  const [error, setError] = useState(undefined);
  const { loadFromSessionStorage, saveToSessionStorage } =
    useSessionStorage(url);
  const [loading, setLoading] = useState(false);
  const [faqs, setFaqs] = useState(loadFromSessionStorage() || []);
  const [currentFAQIndex, setCurrentFAQIndex] = useState(0);

  async function fetchFAQs() {
    try {
      const response = await fetch("/" + url, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();

      data.forEach((d) => {
        d.createdAt = new Date(d.createdAt * 1000).toLocaleString();
      });
      setFaqs(data);
      saveToSessionStorage(data);
    } catch (err) {
      console.error("Error fetching FAQs:", err);
      setError(err);
    }
  }

  function getNextFAQ() {
    if (faqs.length > 0) {
      const nextIndex = Math.floor((currentFAQIndex + 1) % faqs.length);
      setCurrentFAQIndex(nextIndex);
    }
  }

  async function loadFAQs() {
    if (faqs.length === 0) {
      setLoading(true);
      await fetchFAQs();
      setLoading(false);
    }
  }

  useEffect(() => {
    loadFAQs();
  }, []);

  const nextFAQ = faqs[currentFAQIndex];

  return { faqs, error, loading, nextFAQ, getNextFAQ };
}

export default useFAQs;
