// Huffman Coding
#include <list>
#include <iostream>
#include <vector>
#include <string>
#include <memory>


struct Node;
using PNode = std::shared_ptr<Node>;

struct Node {
    explicit Node(char ch, int freq, PNode left=nullptr, PNode right=nullptr): ch(ch), freq(freq), left(left), right(right) {}
    Node(const Node&) = delete;
    Node& operator=(const Node&) = delete;
    ~Node() = default;
    char ch;
    int freq;
    PNode left;
    PNode right;
};

void InsertInSortedList(std::list<PNode>& list, PNode node) {
    std::list<PNode>::iterator it = list.begin();
    while (it != list.end() && (*it)->freq < node->freq) {
        ++it;
    }

    list.insert(it, node);
    for (it = list.begin(); it != list.end(); ++it) {
        std::cout << (*it)->freq << ' ';
    }
    std::cout << "\n";
}


PNode BuildHaffmanCoding(const std::string& st, const std::vector<int>& vec) {
    std::list<PNode> list;
    for (int idx = 0; idx < static_cast<int>(vec.size()); ++idx) {
        InsertInSortedList(list, std::make_shared<Node>(st[idx], vec[idx]));
    }
    while (list.size() > 1) {
        PNode first = list.front();
        list.pop_front();
        PNode second = list.front();
        list.pop_front();
        auto parent = std::make_shared<Node>('.', first->freq + second->freq, first, second);
        InsertInSortedList(list, parent);
    }
    return list.front();
}

void TraverseTree(PNode node, std::string state) {
    if (node->left == nullptr && node->right == nullptr) {
        std::cout << node->ch << ' ' << state << "\n";
        return;
    }
    if (node->left != nullptr) {
        TraverseTree(node->left, state + '0');
    }
    if (node->right != nullptr) {
        TraverseTree(node->right, state + '1');
    }
}

std::string ReadString() {
    std::string str;
    std::cin >> str;
    return str;
}


std::vector<int> ReadVector(int size) {
   std::vector<int> vec(size);
   for (int ii = 0; ii < size; ++ii) {
       std::cin >> vec[ii];
   }
   return vec;
}


int main() {
    std::string str = ReadString();
    std::vector<int> vec = ReadVector(str.length());
    PNode coding = BuildHaffmanCoding(str, vec);
    TraverseTree(coding, "");
}
