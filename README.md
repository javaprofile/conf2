# conf2
**ENHANCING DISTRIBUTED SYSTEMS WITH MULTI-VERSION CONCURRENCY CONTROL OVER 2PL PROTOCOL**
* Author: Vipul Reddy
* Published In : **************************
* Publication Date: 05-2025
* E-ISSN: *************
* Impact Factor: *************
* Link:

**Abstract:**\
This paper addresses performance limitations in database transaction management caused by the blocking nature of the Two-Phase Locking (2PL) protocol. It examines how strict lock-based execution leads to resource contention, longer wait times, and reduced throughput in high-concurrency environments. By highlighting the advantages of Multi-Version Concurrency Control, particularly for read operations, the study shows how MVCC enables non-blocking reads and improves performance while maintaining consistency. The paper emphasizes the need for scalable concurrency control mechanisms that balance correctness and efficiency in modern database systems.

**Key Contributions:**
* **Analysis of Version-Based Storage Growth:**\
Examined how maintaining multiple versions of the same data increases storage usage due to version accumulation, metadata expansion, and cleanup overhead in distributed systems.

* **Use of Optimistic Concurrency Techniques:**\
Applied optimistic concurrency methods as a lightweight approach in low-conflict scenarios to reduce the need for storing historical data versions.
  
* **Comparative Experimental Study:** \
  Conducted a detailed comparison across different cluster sizes, demonstrating that optimistic methods result in lower and more predictable storage growth.
  
* **System Design and Development:**\
  Designed, implemented, and validated storage-efficient concurrency mechanisms using practical experiments in distributed environments.

**Relevance & Real-World Impact**
* **Lower Storage Consumption:**\
Reduced disk usage by limiting unnecessary data version retention while maintaining correct transaction behavior.

* **Improved Scalability:**\
Enabled smoother horizontal scaling of distributed systems by applying optimistic approaches where versioning overhead is unnecessary.

* **Operational Performance Gains:** \
Reduced cleanup overhead and metadata management costs, leading to better throughput and system stability.d metadata management costs, leading to better throughput and system stability.
  
* **Academic and Educational Value:** \
    Provides clear experimental results and implementation insights useful for study and further exploration in concurrency control and distributed system design.

**Experimental Results (Summary)**:

  | Nodes | 2 Phase Locking Protocol | Multi Version Concurrency Control | Reduction (%)   |
  |-------|--------------------------| ----------------------------------| ----------------|
  | 3     |  3000                    | 4000                              | 33.33           |
  | 5     |  4000                    | 6000                              | 50.00           |
  | 7     |  5000                    | 7000                              | 40.00           |
  | 9     |  6000                    | 8000                              | 33.33           |
  | 11    |  7000                    | 9000                              | 28.57           |

**Citation** \
ENHANCING DISTRIBUTED SYSTEMS WITH MULTI-VERSION CONCURRENCY CONTROL OVER 2PL PROTOCOL
* Vipul R 
* Intern***************and Technology 
* E-ISSN *******
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://******
**Author Contact** \
**LinkedIn**: http://linkedin.com/in/Please add here | **Email**: please keep email id @gmail.com







