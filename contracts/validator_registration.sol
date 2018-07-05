<<<<<<< HEAD
pragma solidity ^0.4.23;


contract ValidatorRegistration {
    event ValidatorRegistered(bytes32 pubKey,
    uint256 withdrawalShardID, 
    address withdrawalAddressbytes32, 
    bytes32 randaoCommitment);
=======
pragma solidity 0.4.23;

contract ValidatorRegistration {
    event ValidatorRegistered(
        bytes32 pubKey,
        uint256 withdrawalShardID,
        address withdrawalAddressbytes32,
        bytes32 randaoCommitment
<<<<<<< HEAD
        );
>>>>>>> e7d1cd443b80f1a2f8460d5d4a5026df7926c4b8
=======
    );
>>>>>>> 3769a9eee9c4ad3b51e44be5a35651a840530481

    mapping (bytes32 => bool) public usedPubkey;
    
    uint public constant VALIDATOR_DEPOSIT = 32 ether;

    // Validator registers by sending a transaction of 32ETH to 
    // the following deposit function. The deposit function takes in 
<<<<<<< HEAD
    // validator's public key, withdrawl shard ID (which shard 
    // to send the deposit back to), withdrawl address (which address
    // to send the deposit back to) and randao commitment.
    function deposit(bytes32 _pubkey, 
    uint _withdrawalShardID, 
    address _withdrawalAddressbytes32, 
    bytes32 _randaoCommitment
    ) public payable {
=======
    // validator's public key, withdrawal shard ID (which shard
    // to send the deposit back to), withdrawal address (which address
    // to send the deposit back to) and randao commitment.
    function deposit(
        bytes32 _pubkey,
        uint _withdrawalShardID,
        address _withdrawalAddressbytes32,
        bytes32 _randaoCommitment
        )
        public payable
        {
>>>>>>> e7d1cd443b80f1a2f8460d5d4a5026df7926c4b8
        require(msg.value == VALIDATOR_DEPOSIT);
        require(!usedPubkey[_pubkey]);

        usedPubkey[_pubkey] = true;

        emit ValidatorRegistered(_pubkey, _withdrawalShardID, _withdrawalAddressbytes32, _randaoCommitment);
    }
}