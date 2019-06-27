pragma solidity 0.4.24;

import "./Token.sol";

contract Apply {
    struct addrInfo {
        string finger;
        bool signed;
        bool used;
    }

    mapping (address => addrInfo) addrMap;
    address[] applicants;

    uint256 amount = 66;

    address owner         = 0x0;
    ERC20   token;

    constructor (address _token) public {
        require(_token != 0x0);

        owner = msg.sender;
        token = ERC20(_token);
    }

    function apply(address addr, string finger) external returns(bool) {
        applicants.push(addr);
        addrMap[addr] = addrInfo(finger, false, true);

        return true;
    }

    function sign(address addr, string finger) external returns(bool)  {
        require(addrMap[addr].used, "Address not apply");
        addrMap[addr].finger = finger;

        if (!addrMap[addr].signed) {
            token.transfer(addr, amount);
        }
        addrMap[addr].signed = true;

        return true;
    }

    function getApplicants() external view returns(address[]) {
        return applicants;
    }

    function getSigned() external view returns(address[]) {
        uint256 count;
        address[] memory temp = new address[](applicants.length);
        for (uint256 i = 0; i < applicants.length; i++) {
            address addr = applicants[i];
            if (addrMap[addr].signed) {
                temp[count] = addr;
                count++;
            }
        }

        address[] memory result = new address[](count);
        for (i = 0; i < count; i++) {
            result[i] = temp[i];
        }

        return result;
    }

    function getUnSigned() external view returns(address[]) {
        uint256 count;
        address[] memory temp = new address[](applicants.length);
        for (uint256 i = 0; i < applicants.length; i++) {
            address addr = applicants[i];
            if (addrMap[addr].used && !addrMap[addr].signed) {
                temp[count] = addr;
                count++;
            }
        }

        address[] memory result = new address[](count);
        for (i = 0; i < count; i++) {
            result[i] = temp[i];
        }

        return result;
    }
}
